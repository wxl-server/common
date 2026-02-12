package goretry

import "time"

type config struct {
	retryLimit *int64
	interval   *time.Duration
}

type Option func(*config)
