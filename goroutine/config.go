package goroutine

import "time"

type config struct {
	retryLimit int
	interval   time.Duration
}

type Option func(*config)

func WithRetry(limit int, interval time.Duration) Option {
	return func(c *config) {
		if limit < 0 {
			c.retryLimit = 0
		} else {
			c.retryLimit = limit
		}
		if interval >= 0 {
			c.interval = interval
		}
	}
}
