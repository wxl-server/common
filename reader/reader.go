package reader

import (
	"github.com/qcq1/common/env"
	"github.com/spf13/viper"
)

func InitAppConfig[T struct{}]() *T {
	vip := viper.New()
	vip.SetConfigFile("conf/" + env.GetEnv() + ".yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Config := &T{}
	err = vip.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
	return Config
}
