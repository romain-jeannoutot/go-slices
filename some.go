package goslices

// Some returns `true` if at least one item satisfy specification. Else, return `false`.
func Some[T any](items []T, specification Specification) bool {
	return len(Filter(items, specification)) >= 1
}
