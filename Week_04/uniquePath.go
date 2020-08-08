package week4

// 也可以先处理完边界层才开始遍历
// 也可以用一层数组来遍历
// 还可以直接在多添加一层数组（这个最好）

func uniquePaths(m int, n int) int {
	memo := make([]int, m*n, m*n)
	memo[m*n-1] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {

			if memo[i*m+j] != 0 {
				continue
			}
			bottom, right := 0, 0
			if j+1 < m {
				right = memo[i*m+(j+1)]
			}
			if i+1 < n {
				bottom = memo[(i+1)*m+j]
			}
			memo[i*m+j] = bottom + right
		}
	}
	return memo[0]
}
