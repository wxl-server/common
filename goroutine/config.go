package goroutine

type config struct {
	retryLimit int
}

type Option func(*config)
