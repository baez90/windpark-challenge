package collect

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/baez90/windpark-challenge/sitesim"
)

type StatsPublisher interface {
	Emit(snapshot ParksSnapshot) error
}

type Collector struct {
	Config
	Client    sitesim.Client
	Publisher StatsPublisher
	data      *ParksSnapshot
}

func (c *Collector) Run(ctx context.Context) error {
	logger := zap.L()
	ticker := time.NewTicker(c.CollectInterval)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := c.collect(ctx, time.Now().UTC()); err != nil {
				logger.Error("Failed to collect data", zap.Error(err))
			}
		}
	}

	return nil
}

func (c *Collector) collect(ctx context.Context, timestamp time.Time) error {
	logger := zap.L()
	collectCtx, cancel := context.WithTimeout(ctx, c.CollectInterval/2)
	defer cancel()

	bucket := timestamp.Round(c.BucketSize)
	s, err := c.Client.ListSites(collectCtx)
	if err != nil {
		return err
	}

	if c.data == nil {
		parkSnapshot := &ParksSnapshot{Timestamp: bucket}
		parkSnapshot.Ingest(s)
		c.data = parkSnapshot
	} else if c.data.Timestamp != bucket {
		toEmit := *c.data
		c.data = &ParksSnapshot{Timestamp: bucket}
		c.data.Ingest(s)
		logger.Info("Emit data", zap.Any("data", toEmit))
		if err := c.Publisher.Emit(toEmit); err != nil {
			logger.Error("Failed to publish data", zap.Error(err))
		}
	} else {
		c.data.Ingest(s)
	}

	return nil
}
