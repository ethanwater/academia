package binarytree

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
