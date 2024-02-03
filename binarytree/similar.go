package binarytree

import "reflect"

func (node *TreeNode) SimilarLeaves(node2 *TreeNode) bool {
	var traverse func(*TreeNode, *[]*interface{})
	traverse = func(node *TreeNode, leaves *[]*interface{}) {
		if node == nil {
			return
		}

		if node.IsLeaf() {
			*leaves = append(*leaves, &node.Val)
			return
		}

		traverse(node.Left, leaves)
		traverse(node.Right, leaves)
	}
	var leaves1, leaves2 []*interface{}
	traverse(node, &leaves1)
	traverse(node2, &leaves2)

	return reflect.DeepEqual(leaves1, leaves2)
}
