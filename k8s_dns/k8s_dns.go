package k8s_dns

import (
	"errors"
	"strings"

	"github.com/qcq1/common/choose"
	"github.com/qcq1/common/env"
)

type OptionsStruct struct {
	destServiceName    string
	boeDestServiceName string
}

func WithDestServiceName(destServiceName string) func(*OptionsStruct) {
	split := strings.Split(destServiceName, ".")
	if len(split) != 2 {
		panic(errors.New("destServiceName is not valid, it should be like namespace.serviceName"))
	}
	return func(o *OptionsStruct) {
		o.destServiceName = split[1] + "." + split[0] + ".svc.cluster.local:8888"
	}
}

func WithBoeDestServiceName(boeDestServiceName string) func(*OptionsStruct) {
	return func(o *OptionsStruct) {
		o.boeDestServiceName = boeDestServiceName
	}
}

type Option func(*OptionsStruct)

func BuildK8sDestServiceName(opts ...Option) string {
	options := &OptionsStruct{
		destServiceName:    "",
		boeDestServiceName: "",
	}
	for _, opt := range opts {
		opt(options)
	}

	return choose.If(env.IsProd(), options.destServiceName, options.boeDestServiceName)
}
