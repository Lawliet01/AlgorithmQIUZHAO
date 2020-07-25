var result []string

func generateParenthesis(n int) []string {
	recursive(0, 0, n, "")
	return result
}

func recursive(left int, right int, max int, curr string) {
	if left == max && right == max {
		result = append(result, curr)
		return
	}

	if left < max {
		recursive(left+1, right, max, curr+"(")
	}
	if right < left {
		recursive(left, right+1, max, curr+")")
	}
}