package collect

import "time"

type Config struct {
	CollectInterval time.Duration
	BucketSize      time.Duration
}
