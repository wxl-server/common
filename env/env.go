package env

import "runtime"

type Env = string

const (
	Prod    Env = "prod"
	Boe     Env = "boe"
	Unknown Env = "unknown"
)

func GetEnv() Env {
	sysType := runtime.GOOS

	if sysType == "linux" {
		return Prod
	}

	if sysType == "windows" {
		return Boe
	}
	return Unknown
}

func IsProd() bool {
	return GetEnv() == Prod
}

func IsBoe() bool {
	return GetEnv() == Boe
}
