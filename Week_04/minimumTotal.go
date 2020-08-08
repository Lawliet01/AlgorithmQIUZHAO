package week4

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			left := triangle[i+1][j]
			right := triangle[i+1][j+1]
			if left < right {
				triangle[i][j] = triangle[i][j] + left
			} else {
				triangle[i][j] = triangle[i][j] + right
			}
		}
	}
	return triangle[0][0]
}
