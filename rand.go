package goslices

import (
	"math/rand"
	"time"
)

func Rand[T any](items []T) T {
	if len(items) <= 0 {
		var zero T
		return zero
	}

	rand.Seed(time.Now().UnixNano())
	return items[rand.Intn(len(items))]
}
