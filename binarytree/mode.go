package binarytree

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
