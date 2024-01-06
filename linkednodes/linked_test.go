package linkednodes

import (
	"testing"
)

func TestFetch[T any](t *testing.T) {
	node := &LinkedList[T]{
		Value: *(*T)(1),
		Next: &LinkedList[T]{
			Value: 2,
		},
	}
}
