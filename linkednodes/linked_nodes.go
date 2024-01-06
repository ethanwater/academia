package linkednodes

type LinkedList[T any] struct {
	Value T
	Next  *LinkedList[T]
}

func (node *LinkedList[T]) Fetch() []T {
	if node == nil {
		return nil
	}
	var stack []T
	stack = append(stack, node.Value)
	next := node.Next.Fetch()
	stack = append(stack, next...)

	return stack
}
