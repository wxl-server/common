package goroutine

type config struct {
	retryLimit int
}

type Option func(*config)

func WithRetry(limit int) Option {
	return func(c *config) {
		if limit < 0 {
			c.retryLimit = 0
		} else {
			c.retryLimit = limit
		}
	}
}
