package goroutine

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

// SafeGo 启动Goroutine，panic后会recover
func SafeGo(ctx context.Context, f func(), options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		defer func() {
			if r = recover(); r != nil {
				logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
			}
		}()
		f()
	}()
}

// SafeGoWithParam 启动Goroutine，panic后会recover
func SafeGoWithParam[T any](ctx context.Context, f func(T), arg T, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		defer func() {
			if r = recover(); r != nil {
				logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
			}
		}()
		f(arg)
	}()
}

// MustGo 启动Goroutine，panic后会透传出去
func MustGo(ctx context.Context, f func(), options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		defer func() {
			if r = recover(); r != nil {
				logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
				panic(r)
			}
		}()
		f()
	}()
}

// MustGoWithParam 启动Goroutine，panic后会透传出去
func MustGoWithParam[T any](ctx context.Context, f func(T), arg T, options ...Option) {
	go func() {
		cfg := &config{}
		for _, option := range options {
			option(cfg)
		}
		var r any
		defer func() {
			if r = recover(); r != nil {
				logger.CtxErrorf(ctx, "Goroutine Panic: %v", r)
				panic(r)
			}
		}()
		f(arg)
	}()
}
