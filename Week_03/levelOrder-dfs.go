package week3

func levelOrderDfs(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	dfs(root, 0, &result)

	return result
}

func dfs(node *TreeNode, currLevel int, output *([][]int)) {
	if node == nil {
		return
	}

	for len(*output)-1 < currLevel {
		*output = append(*output, []int{})
	}

	(*output)[currLevel] = append((*output)[currLevel], node.Val)

	dfs(node.Left, currLevel+1, output)
	dfs(node.Right, currLevel+1, output)

}
