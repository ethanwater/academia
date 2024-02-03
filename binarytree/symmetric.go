package binarytree

func (node *TreeNode) IsSymmetric() bool {
	if node == nil {
		return true
	}

	return IsReflection(node.Left, node.Right)
}
