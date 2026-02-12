package goretry

import (
	"time"

	"github.com/wxl-server/common/gptr"
)

func Do(f func(retryTimes int64) error, options ...Option) error {
	cfg := &config{
		retryLimit: gptr.Of(int64(0)),
	}
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
