package goslices

// Includes returns `true` if at least one item in slice is equal to `searchElement`. Else, return `false`.
func Includes[T comparable](items []T, searchElement T) bool {
	for _, item := range items {
		if item == searchElement {
			return true
		}
	}

	return false
}
