package binarytree

func (node *TreeNode) MinimumDepth() int {
	switch {
	case node == nil:
		return 0
	case node.IsLeaf():
		return 1
	default:
		break
	}

	leftDepth := node.Left.MinimumDepth()
	rightDepth := node.Right.MinimumDepth()

	//if either the left or right branches of the node are empty, return
	//the non-empty branch
	if node.Left == nil || node.Right == nil {
		return 1 + max(leftDepth, rightDepth)
	}
	//if left > 0 || right > 0 {
	//return 1 + max(left, right)
	//}

	//return the minimum depth between left and right branches of the node
	return 1 + min(leftDepth, rightDepth)
}
