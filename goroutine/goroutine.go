package goroutine

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
)

// SafeGo 启动Goroutine，panic后会recover
func SafeGo(ctx context.Context, f func() error, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		for retry := 0; retry <= cfg.retryLimit; retry++ {
			err := func() error {
				defer func() {
					if r = recover(); r != nil {
						logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
						panic(r)
					}
				}()
				err := f()
				if err != nil {
					return err
				}
				return nil
			}()
			if r == nil && err == nil {
				break
			}
			time.Sleep(cfg.interval)
		}
	}()
}

// SafeGoWithParam 启动Goroutine，panic后会recover
func SafeGoWithParam[T any](ctx context.Context, f func(T) error, arg T, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		for retry := 0; retry <= cfg.retryLimit; retry++ {
			err := func() error {
				defer func() {
					if r = recover(); r != nil {
						logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
						panic(r)
					}
				}()
				err := f(arg)
				if err != nil {
					return err
				}
				return nil
			}()
			if r == nil && err == nil {
				break
			}
			time.Sleep(cfg.interval)
		}
	}()
}

// MustGo 启动Goroutine，panic后会透传出去
func MustGo(ctx context.Context, f func() error, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		for retry := 0; retry <= cfg.retryLimit; retry++ {
			err := func() error {
				defer func() {
					if r = recover(); r != nil {
						logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
						panic(r)
					}
				}()
				err := f()
				if err != nil {
					return err
				}
				return nil
			}()
			if r == nil && err == nil {
				break
			} else if r != nil && retry == cfg.retryLimit {
				panic(r)
			}
			time.Sleep(cfg.interval)
		}
	}()
}

// MustGoWithParam 启动Goroutine，panic后会透传出去
func MustGoWithParam[T any](ctx context.Context, f func(T) error, arg T, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		for retry := 0; retry <= cfg.retryLimit; retry++ {
			err := func() error {
				defer func() {
					if r = recover(); r != nil {
						logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
						panic(r)
					}
				}()
				err := f(arg)
				if err != nil {
					return err
				}
				return nil
			}()
			if r == nil && err == nil {
				break
			} else if r != nil && retry == cfg.retryLimit {
				panic(r)
			}
			time.Sleep(cfg.interval)
		}
	}()
}
