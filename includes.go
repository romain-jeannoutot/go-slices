package goslices

func Includes[T comparable](items []T, searchElement T) bool {
	for _, item := range items {
		if item == searchElement {
			return true
		}
	}

	return false
}
