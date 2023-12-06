package binarytree

import (
	"testing"
)

func TestMinimumDepth(t *testing.T) {
	//  	 1
	//  	/ \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	//  /    \
	//  8     9
	//  	  /
	//  	 10

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 10,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	answer := 3
	if root.MinimumDepth() != int32(answer) {
		t.Errorf("result=%d", root.MinimumDepth())
	}
}

func TestNodeDepthInt(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 10,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	answer := 4
	if root.NodeDepthInt(9) != answer {
		t.Errorf("result=%d", root.NodeDepthInt(9))
	}
}
