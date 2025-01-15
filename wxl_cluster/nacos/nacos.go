package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/qcq1/common/choose"
	"github.com/qcq1/common/env"
)

func NewNacosClient() (naming_client.INamingClient, error) {
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
}
