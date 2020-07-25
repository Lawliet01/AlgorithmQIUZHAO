package week2

func preorderTraversal(root *TreeNode) []int {
	list := []int{}
	preHelper(root, &list)
	return list
}

func preHelper(root *TreeNode, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		if root.Left != nil {
			preHelper(root.Left, list)
		}
		if root.Right != nil {
			preHelper(root.Right, list)
		}
	}
}
