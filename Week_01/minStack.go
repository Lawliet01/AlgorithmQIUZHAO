import "math"

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{math.MaxInt64},
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	topMin := this.min[len(this.min)-1]
	this.min = append(this.min, min(topMin, x))
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}