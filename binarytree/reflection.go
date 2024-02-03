package binarytree

func IsReflection(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}

	return x.Val == y.Val && IsReflection(x.Left, y.Right) && IsReflection(x.Right, y.Left)
}
