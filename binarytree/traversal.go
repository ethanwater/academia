package binarytree

func (node *TreeNode) InorderTraversal() {
	if node == nil {
		return
	}

	stack := []*TreeNode{}
	currentNode := node

	for currentNode != nil || len(stack) > 0 {
		if currentNode != nil {
			stack = append(stack, currentNode)
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentNode = currentNode.Right
	}
}
