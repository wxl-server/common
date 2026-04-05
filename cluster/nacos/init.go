package nacos

type NacosConfig struct {
	Username *string
	Password *string
	Prod     *SubNacosConfig
	Boe      *SubNacosConfig
}

type SubNacosConfig struct {
	NamespaceId *string
	IP          *string
	Port        *int64
}

var nacosConfig *NacosConfig

func Init(nacosConfig1 NacosConfig) {
	nacosConfig = &nacosConfig1
}
