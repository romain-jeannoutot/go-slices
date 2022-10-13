package goslices

func Find[T any](items []T, specification Specification) T {
	for _, item := range items {
		if specification.IsSatisfied(item) {
			return item
		}
	}

	var zero T
	return zero
}
