package goslices

func Some[T any](items []T, specification Specification) bool {
	return len(Filter(items, specification)) >= 1
}
