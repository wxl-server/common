package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wxl-server/common/choose"
	"github.com/wxl-server/common/env"
)

func NewNacosClient() (naming_client.INamingClient, error) {
	if env.IsProd() {
		return clients.NewNamingClient(
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
	} else {
		return clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{
					NamespaceId: "public",
					Username:    "nacos",
					Password:    "wxl5211314",
				},
				ServerConfigs: []constant.ServerConfig{
					*constant.NewServerConfig("127.0.0.1", 8848),
				},
			},
		)
	}
}
