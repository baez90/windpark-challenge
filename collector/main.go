package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"

	"github.com/baez90/windpark-challenge/internal/collect"
	"github.com/baez90/windpark-challenge/internal/publishing"
	"github.com/baez90/windpark-challenge/sitesim"
)

var (
	logger         *zap.Logger
	appCtx         context.Context
	appStop        context.CancelFunc
	siteSimBaseURL string
	amqpUri        string
	amqpCfg        amqp091.Config
	routingKeys    []string
	collectCfg     collect.Config
	rootCmd        = &cobra.Command{
		Use: "windpark",
		PersistentPreRunE: func(*cobra.Command, []string) (err error) {
			if logger, err = zap.NewProduction(); err != nil {
				return err
			}

			zap.ReplaceGlobals(logger)

			appCtx, appStop = signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

			return
		},
	}
	collectCmd = &cobra.Command{
		Use: "collect",
		RunE: func(*cobra.Command, []string) error {
			publisher, err := rabbitmq.NewPublisher(amqpUri, amqpCfg, rabbitmq.WithPublisherOptionsLogging)
			if err != nil {
				return err
			}

			retryClient := retryablehttp.NewClient()
			retryClient.Logger = nil
			retryClient.RetryMax = 10

			collector := collect.Collector{
				Config: collectCfg,
				Publisher: &publishing.RabbitMQPublisher{
					Publisher:   publisher,
					RoutingKeys: routingKeys,
				},
				Client: sitesim.Client{
					Client:  retryClient.StandardClient(),
					BaseURL: siteSimBaseURL,
				},
			}

			return collector.Run(appCtx)
		},
	}
)

func main() {
	collectCmd.Flags().StringVar(&siteSimBaseURL, "base-url", "http://renewables-codechallenge.azurewebsites.net", "Base URL for API calls")
	collectCmd.Flags().StringVar(&amqpUri, "publishing-uri", "amqp://rabbitmq:rabbitmq@localhost:5672/", "URI to connect to RabbitMQ")
	collectCmd.Flags().StringVar(&amqpCfg.Vhost, "publish-vhost", "/", "RabbitMQ vHost to use")
	collectCmd.Flags().StringSliceVar(&routingKeys, "publishing-routing-keys", []string{"windpark-challenge"}, "RabbitMQ vHost to use")
	collectCmd.Flags().DurationVar(&collectCfg.CollectInterval, "collect-interval", 5*time.Second, "Interval how often data is being collected")
	collectCmd.Flags().DurationVar(&collectCfg.BucketSize, "bucket-size", 5*time.Minute, "Interval how often data is being collected")
	rootCmd.AddCommand(collectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error during execution: %v\n", err)
		os.Exit(1)
	}
}
