package training

func moveZeroes(nums []int) {
	curr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[curr] = nums[i]
			curr++
		}
	}

	for i := curr; i < len(nums); i++ {
		nums[i] = 0
	}
}
