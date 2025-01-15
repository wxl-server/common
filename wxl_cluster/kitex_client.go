package wxl_cluster

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/qcq1/common/gptr"
	"github.com/qcq1/common/wxl_cluster/kitex_middleware"
	"github.com/qcq1/common/wxl_cluster/nacos"
)

const (
	MaxRetryTimes = 3
)

type RawCallInterface[T any] interface {
	setClient(T)
	setRetryPolicy(*retry.FailurePolicy)
}

type newClientFunc[T any] func(destService string, opts ...client.Option) (T, error)

func NewClient[T any, M any](newClient newClientFunc[T], destService string, RawCall *RawCallInterface[T]) {
	cli, err := nacos.NewNacosClient()
	if err != nil {
		logger.Errorf("[Init] clients.NewNamingClient failed, err = %v", err)
		panic(err)
	}
	destClient, err := newClient(
		destService,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMiddleware(kitex_middleware.LogMiddleware),
	)
	gptr.Indirect(RawCall).setClient(destClient)
	if err != nil {
		logger.Errorf("[Init] minercore.NewClient failed, err = %v", err)
		panic(err)
	}
	rp := retry.NewFailurePolicy()
	rp.WithMaxRetryTimes(MaxRetryTimes)
	gptr.Indirect(RawCall).setRetryPolicy(rp)
}
