package binarytree

func (node *TreeNode) MaximumDepth() int {
	switch {
	case node == nil:
		return 0
	case node.IsLeaf():
		return 1
	}

	left := node.Left.MaximumDepth()
	right := node.Right.MaximumDepth()

	if node.Left == nil || node.Right == nil {
		return 1 + max(left, right)
	}

	return 1 + max(left, right)
}
