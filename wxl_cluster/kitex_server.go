package wxl_cluster

import (
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/qcq1/common/choose"
	"github.com/qcq1/common/env"
	"github.com/qcq1/common/wxl_cluster/kitex_middleware"
)

type newServerFunc[T any] func(handler T, opts ...server.Option) server.Server

func NewServer[T any](newServer newServerFunc[T], handler T, serverName string) {
	options := make([]server.Option, 0)
	// boe 环境下指定服务地址
	if env.IsBoe() {
		host := ""
		port := 8888
		retry := 0
		for {
			addr, err := net.ResolveTCPAddr("tcp", buildHostPort(host, port))
			if err != nil && retry >= 10 {
				panic(err)
			} else if err != nil {
				port += 1
				retry += 1
			} else {
				options = append(options, server.WithServiceAddr(addr))
				break
			}
		}
	}

	// 注册服务到nacos
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: choose.If(env.IsProd(), "public", "boe"),
				Username:    "nacos",
				Password:    "wxl5211314",
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig("wxl475.cn", 30898),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	options = append(options, server.WithRegistry(registry.NewNacosRegistry(cli)))
	options = append(options, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serverName}))

	// 日志中间件
	options = append(options, server.WithMiddleware(kitex_middleware.LogMiddleware))

	// 启动服务
	svr := newServer(handler, options...)
	if err = svr.Run(); err != nil {
		panic(err)
	}
}

func buildHostPort(host string, port int) string {
	return host + ":" + strconv.FormatInt(int64(port), 10)
}
