package goslices

func Every[T any](items []T, specification Specification) bool {
	return len(Filter(items, specification)) == len(items)
}
