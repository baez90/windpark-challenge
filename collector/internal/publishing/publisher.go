package publishing

import (
	"encoding/json"

	"github.com/wagslane/go-rabbitmq"

	"github.com/baez90/windpark-challenge/internal/collect"
)

type RabbitMQPublisher struct {
	Publisher   *rabbitmq.Publisher
	RoutingKeys []string
}

func (p *RabbitMQPublisher) Emit(snapshot collect.ParksSnapshot) error {
	data, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}
	err = p.Publisher.Publish(
		data,
		p.RoutingKeys,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsExchange("amq.topic"),
	)

	if err != nil {
		return err
	}
	return nil
}
