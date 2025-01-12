package gslice

func ForEach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

func Map[T, U any](s []T, f func(T) U) (r []U) {
	r = make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return
}
