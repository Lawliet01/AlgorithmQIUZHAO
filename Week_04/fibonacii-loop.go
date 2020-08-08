package week4

func fibonacci(num int) int {
	memo := make([]int, num+1, num+1)
	memo[0] = 0
	memo[1] = 1
	i := 2
	for i <= num {
		memo[i] = memo[i-1] + memo[i-2]
		i++
	}
	return memo[num]
}
