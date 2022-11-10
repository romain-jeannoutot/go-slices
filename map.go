package goslices

// Map creates a new slice populated with elements returned by function called on each slice element.
func Map[T any, S any](items []T, callback func(item T, idx int, items []T) S) []S {
	ss := make([]S, 0)
	for idx, item := range items {
		ss = append(ss, callback(item, idx, items))
	}

	return ss
}
