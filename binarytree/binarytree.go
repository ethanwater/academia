package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) MinimumDepth() int32 {
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

	//return the minimum depth between left and right branches of the node
	return 1 + min(leftDepth, rightDepth)
}

func (node *TreeNode) NodeDepthInt(target int) int {
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

	leftNode := node.Left.NodeDepthInt(target)
	if leftNode > 0 {
		return 1 + leftNode
	}

	rightNode := node.Right.NodeDepthInt(target)
	if rightNode > 0 {
		return 1 + rightNode
	}

	return 0 //if no matching node is detected
}

// leaf node: a node with no children
func (node *TreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}
