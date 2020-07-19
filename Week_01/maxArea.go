package training

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		var tempArea int
		if height[i] > height[j] {
			tempArea = (j - i) * height[j]
			j--
		} else {
			tempArea = (j - i) * height[i]
			i++
		}
		if tempArea > max {
			max = tempArea
		}
	}
	return max
}
