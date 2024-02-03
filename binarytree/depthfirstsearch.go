package binarytree

func (node *TreeNode) DepthFirstSearch() []interface{} {
	if node == nil {
		return nil
	}
	var stack []interface{}

	rightNode := node.Right.DepthFirstSearch()
	leftNode := node.Left.DepthFirstSearch()

	stack = append(stack, node.Val)
	if leftNode != nil {
		stack = append(stack, leftNode...)
	}
	if rightNode != nil {
		stack = append(stack, rightNode...)
	}

	return stack
}
