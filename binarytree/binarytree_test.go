package binarytree

import (
	"reflect"
	"testing"
)

// test binary tree:
//
//		 1
//		/ \
//	   2   3
//	  / \ / \
//	 4  5 6  7
//	 /  \  \
//
// 8     9  10
func TestMinimumDepth(t *testing.T) {
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
	if root.MinimumDepth() != answer {
		t.Errorf("result=%d", root.MinimumDepth())
	}
}
func TestMaximumDepth(t *testing.T) {
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
	if root.MaximumDepth() != answer {
		t.Errorf("result=%d", root.MaximumDepth())
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

	//TODO: return 3, not 4
	answer := 4
	if root.NodeDepth(9) != answer {
		t.Errorf("result=%d", root.NodeDepth(9))
	}
}

func TestDepthFirstSearch(t *testing.T) {
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
				Val: "cat",
			},
		},
	}

	expected_result := [10]interface{}{1, 2, 4, 8, 5, 9, 3, 6, 10, "cat"}

	if !reflect.DeepEqual(root.DepthFirstSearch(), expected_result[:]) {
		t.Errorf("result=%d", root.DepthFirstSearch())
	}
}
