package training

func climbStairs1(n int) int {
	memo := make([]int, n, n)
	return climbEach1(n, 0, memo)
}

func climbEach1(n, i int, memo []int) int {
	if i == n {
		return 1
	}

	if i > n {
		return 0
	}
	if memo[i] == 0 {
		memo[i] = climbEach1(n, i+1, memo) + climbEach1(n, i+2, memo)
	}
	return memo[i]
}
