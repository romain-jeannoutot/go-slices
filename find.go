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

func FindIndex[T any](items []T, specification Specification) int {
	for idx, item := range items {
		if specification.IsSatisfied(item) {
			return idx
		}
	}

	return -1
}
