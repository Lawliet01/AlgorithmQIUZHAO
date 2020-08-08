package week4

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	lastPositive := 0
	lastNegative := 0
	max := 0

	for i := 0; i < len(nums); i++ {
		store1 := nums[i]
		store2 := nums[i] * lastPositive
		store3 := nums[i] * lastNegative

		lastPositive, lastNegative = findMaxAndMin([]int{store1, store2, store3})

		if lastPositive > max {
			max = lastPositive
		}
	}

	return max
}

func findMaxAndMin(nums []int) (max, min int) {
	for _, v := range nums {
		if v > max {
			max = v
			continue
		}

		if v < min {
			min = v
			continue
		}
	}

	return
}
