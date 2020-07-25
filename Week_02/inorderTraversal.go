package week2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	list := []int{}
	list = helper(root, list)
	return list
}

func helper(root *TreeNode, list []int) []int {
	if root != nil {
		if root.Left != nil {
			list = helper(root.Left, list)
		}
		list = append(list, root.Val)
		if root.Right != nil {
			list = helper(root.Right, list)
		}
	}
	return list
}
