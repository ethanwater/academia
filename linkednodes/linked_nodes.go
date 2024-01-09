package linkednodes

type LinkedList[T any] struct {
	Value T
	Next  *LinkedList[T]
}

type LinkedNode struct {
	Val int
	Next *LinkedNode
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

func (node *LinkedNode) Merge(node2 *LinkedNode) *LinkedNode {
	sentinel_node := &LinkedNode{}
	merging_node := sentinel_node

	switch{
	case node2 == nil:
		return node
	case node == nil:
		return node2
	}

	for node!=nil && node2!=nil{
		switch{
		case node.Val > node2.Val:
			merging_node.Next = node
			node = node.Next
		case node.Val < node2.Val:
			merging_node.Next = node2
			node2 = node2.Next
		}
		merging_node = merging_node.Next
	}

	switch{
	case node == nil && node2 != nil:
		merging_node.Next = node2
	case node != nil && node2 == nil:
		merging_node.Next = node
	}

	return sentinel_node
}

func (node *LinkedNode) InplaceMerge(node2 *LinkedNode) {

}

