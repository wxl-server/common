package nacos

import (
	"errors"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wxl-server/common/env"
)

func NewNacosClient() (naming_client.INamingClient, error) {
	if env.IsProd() {
		return newProdNacosClient()
	} else {
		client, err := newProdNacosClient()
		if err != nil {
			return newBoeNacosClient()
		}
		return client, nil
	}
}

func newProdNacosClient() (naming_client.INamingClient, error) {
	if nacosConfig == nil ||
		nacosConfig.Prod == nil ||
		nacosConfig.Prod.NamespaceId == nil ||
		nacosConfig.Prod.IP == nil ||
		nacosConfig.Prod.Port == nil ||
		nacosConfig.Username == nil ||
		nacosConfig.Password == nil {
		return nil, errors.New("nacos config not init")
	}
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: *nacosConfig.Prod.NamespaceId,
				Username:    *nacosConfig.Username,
				Password:    *nacosConfig.Password,
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(*nacosConfig.Prod.IP, uint64(*nacosConfig.Prod.Port)),
			},
		},
	)
}

func newBoeNacosClient() (naming_client.INamingClient, error) {
	if nacosConfig == nil ||
		nacosConfig.Boe == nil ||
		nacosConfig.Boe.NamespaceId == nil ||
		nacosConfig.Boe.IP == nil ||
		nacosConfig.Boe.Port == nil ||
		nacosConfig.Username == nil ||
		nacosConfig.Password == nil {
		return nil, errors.New("nacos config not init")
	}
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: *nacosConfig.Boe.NamespaceId,
				Username:    *nacosConfig.Username,
				Password:    *nacosConfig.Password,
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(*nacosConfig.Boe.IP, uint64(*nacosConfig.Boe.Port)),
			},
		},
	)
}
