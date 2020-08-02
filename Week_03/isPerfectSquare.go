package week3

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}

	left := 0
	right := num

	for left <= right {
		mid := (left + right) / 2
		res := mid * mid
		if res == num {
			return true
		} else if res < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
