package binarytree

func IsIdentical(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}

	return x.Val == y.Val && IsIdentical(x.Left, y.Left) && IsIdentical(x.Right, y.Right)
}
