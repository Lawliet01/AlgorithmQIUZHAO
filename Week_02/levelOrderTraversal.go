package week2

func levelOrder(root *Node) [][]int {
	var queue []*Node
	var res [][]int

	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}

		for _, child := range queue {

			queue = queue[1:]

			if child != nil {
				level = append(level, child.Val)
				queue = append(queue, child.Children...)
			}

		}

		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}
