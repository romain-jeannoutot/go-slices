package goslices

import (
	"math/rand"
	"time"
)

// Rand returns a random element from slice. Empty value if slice is empty.
func Rand[T any](items []T) T {
	if len(items) <= 0 {
		var zero T
		return zero
	}

	rand.Seed(time.Now().UnixNano())
	return items[rand.Intn(len(items))]
}
