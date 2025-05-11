package internal

import (
	"iter"
)

func Map[T any, R Renderer](items iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range items {
			node := transform(item)
			if !yield(node) {
				return
			}
		}
	}
}
