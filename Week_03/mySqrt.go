package week3

func mySqrt(x int) int {
	left := 0
	right := x
	var res int
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			res = mid
			break
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
