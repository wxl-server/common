package goretry

import "time"

type config struct {
	retryLimit *int64
	interval   *time.Duration
}

type Option func(*config)

func WithRetryLimit(retryLimit int64) Option {
	return func(cfg *config) {
		cfg.retryLimit = &retryLimit
	}
}
func WithInterval(interval time.Duration) Option {
	return func(cfg *config) {
		cfg.interval = &interval
	}
}
