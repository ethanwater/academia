package binarytree

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
