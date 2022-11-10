package goslices

// Find returns the first item in slice satisfying specification. Else, return empty value.
func Find[T any](items []T, specification Specification) T {
	for _, item := range items {
		if specification.IsSatisfied(item) {
			return item
		}
	}

	var zero T
	return zero
}

// FindIndex returns index of the first item in slice satisfying specification. Else, return -1.
func FindIndex[T any](items []T, specification Specification) int {
	for idx, item := range items {
		if specification.IsSatisfied(item) {
			return idx
		}
	}

	return -1
}
