package env

import "runtime"

const (
	prod    = "prod"
	boe     = "boe"
	unknown = "unknown"
)

func GetEnv() string {
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
