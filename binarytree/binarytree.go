package binarytree

import (
	"reflect"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

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

func (node *TreeNode) MaximumDepth() int {
	switch {
	case node == nil:
		return 0
	case node.IsLeaf():
		return 1
	}

	left := node.Left.MaximumDepth()
	right := node.Right.MaximumDepth()

	if node.Left == nil || node.Right == nil {
		return 1 + max(left, right)
	}

	return 1 + max(left, right)
}

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

func (node *TreeNode) DepthFirstSearch() []interface{} {
	if node == nil {
		return nil
	}
	var stack []interface{}

	rightNode := node.Right.DepthFirstSearch()
	leftNode := node.Left.DepthFirstSearch()

	stack = append(stack, node.Val)
	if leftNode != nil {
		stack = append(stack, leftNode...)
	}
	if rightNode != nil {
		stack = append(stack, rightNode...)
	}

	return stack
}

func (node *TreeNode) IsSymmetric() bool {
	if node == nil {
		return true
	}

	return isReflection(node.Left, node.Right)
}

func (node *TreeNode) IsBalanced() bool {
	if node == nil {
		return true
	}

	var heightCalc func(*TreeNode) int
	heightCalc = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := heightCalc(node.Left)
		right := heightCalc(node.Right)

		if left == -1 || right == -1 || max(left, right)-min(left, right) > 1 {
			return -1
		}

		return 1 + max(left, right)
	}

	return heightCalc(node) != -1
}

func isReflection(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}

	return x.Val == y.Val && isReflection(x.Left, y.Right) && isReflection(x.Right, y.Left)
}

func IsIdentical(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}

	return x.Val == y.Val && IsIdentical(x.Left, y.Left) && IsIdentical(x.Right, y.Right)
}

func (node *TreeNode) FindAllPaths() [][]interface{} {
	var paths [][]interface{}
	var currentPath []interface{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		currentPath = append(currentPath, node.Val)

		if node.IsLeaf() {
			path := make([]interface{}, len(currentPath))
			copy(path, currentPath)
			paths = append(paths, path)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}

		currentPath = currentPath[:len(currentPath)-1] // Backtrack
	}

	dfs(node)
	return paths
}

func (node *TreeNode) Mode() []int {
	if node == nil {
		return nil
	}
	var (
		mode     int
		parent   interface{}
		count    int
		modes    []int
		traverse func(*TreeNode)
		contains func([]int, int) bool
	)

	contains = func(container []int, value int) bool {
		for _, v := range container {
			if v == value {
				return true
			}
		}
		return false
	}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		if node.Val == parent {
			count++
		} else {
			count = 1
			parent = node.Val
		}

		if count > mode {
			mode = count
			modes = make([]int, mode)
		}
		if count == mode && !contains(modes, mode) {
			modes = append(modes, mode)
		}
		traverse(node.Right)
	}

	traverse(node)
	return modes
}

func (node *TreeNode) InorderTraversal() {
	if node == nil {
		return
	}

	stack := []*TreeNode{}
	currentNode := node

	for currentNode != nil || len(stack) > 0 {
		if currentNode != nil {
			stack = append(stack, currentNode)
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentNode = currentNode.Right
	}
}

func (node *TreeNode) InorderTraversal2() []*TreeNode {
	if node == nil {
		return nil
	}

	var stack []*TreeNode
	stack = append(stack, node.Left.InorderTraversal2()...)
	stack = append(stack, node)
	stack = append(stack, node.Right.InorderTraversal2()...)
	return stack
}

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

//func (root *TreeNode) HasPathSum(target int) {
//	  if root == nil {
//        return false
//    }
//    if root.IsLeaf {
//        return targetSum == root.Val
//    }
//    return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
//}
