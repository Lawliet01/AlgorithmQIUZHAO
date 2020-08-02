package week3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currLen := len(queue)
		currLevel := []int{}
		for i := 0; i < currLen; i++ {
			n := queue[i]
			if n == nil {
				continue
			}
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
			currLevel = append(currLevel, n.Val)
		}

		queue = queue[currLen:]

		if len(queue) != 0 {
			result = append(result, currLevel)
		}
	}
	return result
}
