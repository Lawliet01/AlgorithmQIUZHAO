package week4

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	res := 0
	last1 := 1
	last2 := 2

	for i := 3; i < n+1; i++ {
		res = last1 + last2
		last1 = last2
		last2 = res
	}
	return res
}
