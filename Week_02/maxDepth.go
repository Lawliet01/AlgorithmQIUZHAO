func maxDepth(root *TreeNode) int {
	var max int = 0
	if root != nil {
		max = addDepth(root, max)
	}
	return max
}

func addDepth(node *TreeNode, currMax int) int {
	if node.Left == nil && node.Right == nil {
		return currMax + 1
	}

	leftMax, rightMax := currMax, currMax

	if node.Left != nil {
		leftMax = addDepth(node.Left, currMax+1)
	}

	if node.Right != nil {
		rightMax = addDepth(node.Right, currMax+1)
	}

	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}