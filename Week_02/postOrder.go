package week2

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	postHelper(root, &res)
	return res
}

func postHelper(node *Node, list *[]int) {
	if node != nil {
		for _, n := range node.Children {
			postHelper(n, list)
		}
		*list = append(*list, node.Val)
	}
}
