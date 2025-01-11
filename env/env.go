package env

import "runtime"

type Env = string

const (
	prod    Env = "prod"
	boe     Env = "boe"
	unknown Env = "unknown"
)

func GetEnv() Env {
	sysType := runtime.GOOS

	if sysType == "linux" {
		return prod
	}

	if sysType == "windows" {
		return boe
	}
	return unknown
}

func IsProd() bool {
	return GetEnv() == prod
}

func IsBoe() bool {
	return GetEnv() == boe
}
