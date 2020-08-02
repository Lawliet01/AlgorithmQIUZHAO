package week3

func canJump(nums []int) bool {
	endReachable := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= endReachable {
			endReachable = i
		}
	}
	return endReachable == 0
}
