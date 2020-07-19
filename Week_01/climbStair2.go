package training

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var curr int
	prev := 2
	prevPrev := 1
	for i := 3; i <= n; i++ {
		curr = prev + prevPrev
		prevPrev = prev
		prev = curr
	}
	return curr
}
