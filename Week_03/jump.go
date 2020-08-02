package week3

func jump(nums []int) int {
	end := len(nums) - 1
	res := 0
	for {
		for i := 0; i < end; i++ {
			if nums[i]+i >= end {
				end = i
				res++
				break
			}
		}

		if end == 0 {
			return res
		}
	}
}
