package week4

// 不单独处理边界情况，因为要写很多循环

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width := len(obstacleGrid[0])
	height := len(obstacleGrid)
	memo := make([]int, width, width)

	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {

			bottom, right := 0, 0

			if i == height-1 && j == width-1 && obstacleGrid[i][j] == 0 {
				memo[j] = 1
				continue
			}

			if j+1 < width {
				right = memo[j+1]
			}

			if i+1 < height {
				bottom = memo[j]
			}

			if obstacleGrid[i][j] == 1 {
				memo[j] = 0
			} else {
				memo[j] = right + bottom
			}
		}
	}

	return memo[0]
}
