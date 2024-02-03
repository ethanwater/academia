package binarytree

func (node *TreeNode) NodeDepth(target interface{}) int {
	switch {
	case node == nil:
		return 0
	case node.Val == target:
		return 1
	case node.IsLeaf():
		return 0
	default:
		break
	}

	leftNode := node.Left.NodeDepth(target)
	if leftNode > 0 {
		return 1 + leftNode
	}

	rightNode := node.Right.NodeDepth(target)
	if rightNode > 0 {
		return 1 + rightNode
	}

	return 0 //if no matching node is detected
}
