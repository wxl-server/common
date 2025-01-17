package gptr

import "github.com/wxl-server/common/gvalue"

func Of[T any](v T) *T {
	return &v
}

func Indirect[T any](p *T) T {
	if p == nil {
		return gvalue.Zero[T]()
	}
	return *p
}
