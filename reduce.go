package goslices

func Reduce[T any, S any](items []T, callback func(acc S, item T, idx int, items []T) S, initialValue S) S {
	s := initialValue
	for idx, item := range items {
		s = callback(s, item, idx, items)
	}

	return s
}
