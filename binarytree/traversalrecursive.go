package binarytree

func (node *TreeNode) InorderTraversal2() []*TreeNode {
	if node == nil {
		return nil
	}

	var stack []*TreeNode
	stack = append(stack, node.Left.InorderTraversal2()...)
	stack = append(stack, node)
	stack = append(stack, node.Right.InorderTraversal2()...)
	return stack
}
