package goslices

func Map[T any, S any](items []T, callback func(item T) S) []S {
	ss := make([]S, 0)
	for _, item := range items {
		ss = append(ss, callback(item))
	}

	return ss
}
