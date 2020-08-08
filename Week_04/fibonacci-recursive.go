package week4

func fibo(num int) int {
	memo := make([]int, num+1, num+1)
	res := helper(num, &memo)
	return res
}

func helper(n int, memo *([]int)) int {
	if n <= 1 {
		return n
	}

	if (*memo)[n] == 0 {
		(*memo)[n] = helper(n-1, memo) + helper(n-2, memo)
	}

	return (*memo)[n]
}
