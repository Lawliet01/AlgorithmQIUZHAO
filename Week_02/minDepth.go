func minDepth(root *TreeNode) int {
	min := 0
	if root != nil {
		min = findMinDepth(root, min+1)
	}
	return min
}

func findMinDepth(node *TreeNode, min int) int {
	if node.Left == nil && node.Right == nil {
		return min
	}

	if node.Left == nil && node.Right != nil {
		return findMinDepth(node.Right, min+1)
	}

	if node.Left != nil && node.Right == nil {
		return findMinDepth(node.Left, min+1)
	}

	leftMin := findMinDepth(node.Left, min+1)
	rightMin := findMinDepth(node.Right, min+1)

	if leftMin > rightMin {
		return rightMin
	}
	return leftMin
}