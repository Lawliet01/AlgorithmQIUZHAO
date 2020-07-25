var testResult bool = true

func isValidBST(root *TreeNode) bool {
	if root != nil {
		testBST(root)
	}

	return testResult
}

func testBST(node *TreeNode) {
	if !testResult {
		return
	}

	if node.Left != nil && node.Left.Val < node.Val {
		testBST(node.Left)
	} else if node.Left != nil && node.Left.Val > node.Val {
		testResult = false
		return
	}

	if node.Right != nil && node.Right.Val > node.Val {
		testBST(node.Right)
	} else if node.Right != nil && node.Right.Val < node.Val {
		testResult = false
		return
	}
}