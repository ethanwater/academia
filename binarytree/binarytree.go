package binarytree

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}
