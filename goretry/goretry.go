package goretry

import (
	"context"
	"time"

	"github.com/wxl-server/common/gptr"
)

func Do(ctx context.Context, f func(retryTimes int64) error, options ...Option) error {
	cfg := &config{}
	for _, option := range options {
		option(cfg)
	}
	var err error
	for retry := int64(0); retry <= gptr.Indirect(cfg.retryLimit); retry++ {
		err = f(retry)
		if err != nil {
			if cfg.interval != nil {
				time.Sleep(*cfg.interval)
			}
			continue
		}
		return nil
	}
	return err
}
