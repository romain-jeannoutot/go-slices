package goslices

// Every returns `true` if all items satisfy specification. Else, return `false`.
func Every[T any](items []T, specification Specification) bool {
	return len(Filter(items, specification)) == len(items)
}
