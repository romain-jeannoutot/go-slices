package goslices

func Filter[T any](items []T, specification Specification) []T {
	ss := make([]T, 0)
	for _, item := range items {
		if specification.IsSatisfied(item) {
			ss = append(ss, item)
		}
	}

	return ss
}
