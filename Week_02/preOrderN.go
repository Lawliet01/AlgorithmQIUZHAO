package week2

func preorderN(root *Node) []int {
	res := []int{}
	preOrderHelper(root, &res)
	return res
}

func preOrderHelper(node *Node, list *[]int) {
	if node != nil {
		*list = append(*list, node.Val)
		for _, n := range node.Children {
			preOrderHelper(n, list)
		}
	}
}
