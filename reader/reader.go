package reader

import (
	"github.com/spf13/viper"
	"github.com/wxl-server/common/env"
)

func InitAppConfig[T any]() *T {
	vip := viper.New()
	vip.SetConfigFile("conf/" + env.GetEnv() + ".yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Config := new(T)
	err = vip.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
	return Config
}
