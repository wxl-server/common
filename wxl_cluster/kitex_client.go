package wxl_cluster

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/qcq1/common/wxl_cluster/kitex_middleware"
	"github.com/qcq1/common/wxl_cluster/nacos"
)

const (
	MaxRetryTimes = 3
)

type RawCallStruct[T any] struct {
	client T
	rp     *retry.FailurePolicy
}

type newClientFunc[T any] func(destService string, opts ...client.Option) (T, error)

func NewClient[T any](newClient newClientFunc[T], destService string) *RawCallStruct[T] {
	RawCall := &RawCallStruct[T]{}

	cli, err := nacos.NewNacosClient()
	if err != nil {
		logger.Errorf("[Init] clients.NewNamingClient failed, err = %v", err)
		panic(err)
	}
	RawCall.client, err = newClient(
		destService,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMiddleware(kitex_middleware.LogMiddleware),
	)
	if err != nil {
		logger.Errorf("[Init] minercore.NewClient failed, err = %v", err)
		panic(err)
	}
	rp := retry.NewFailurePolicy()
	rp.WithMaxRetryTimes(MaxRetryTimes)
	RawCall.rp = rp
	return RawCall
}
