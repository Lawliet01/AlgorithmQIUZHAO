package week2

type NumberFrequent struct {
	num       int
	frequency int
}

type MaxHeap struct {
	data []*NumberFrequent
}

func InitMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []*NumberFrequent{},
	}
}

func (this *MaxHeap) insert(n *NumberFrequent) {
	this.data = append(this.data, n)
	currIndex := len(this.data) - 1
	for {
		parentNodeIndex := (currIndex - 1) / 2

		if this.data[parentNodeIndex].frequency > this.data[currIndex].frequency || currIndex == 0 {
			return
		}

		this.data[parentNodeIndex], this.data[currIndex] = this.data[currIndex], this.data[parentNodeIndex]
		currIndex = parentNodeIndex
	}
}

func (this *MaxHeap) pop() (result *NumberFrequent) {
	result = this.data[0]
	this.data[0] = this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	length := len(this.data)
	currIndex := 0
	for {

		leftChild := 2*currIndex + 1
		rightChild := 2*currIndex + 2

		if leftChild >= length {
			break
		} else if leftChild < length && rightChild >= length {
			if this.data[leftChild].frequency > this.data[currIndex].frequency {
				this.data[currIndex], this.data[leftChild] = this.data[leftChild], this.data[currIndex]
			}
			break
		}

		if this.data[currIndex].frequency < this.data[leftChild].frequency || this.data[currIndex].frequency < this.data[rightChild].frequency {
			if this.data[leftChild].frequency > this.data[rightChild].frequency {
				this.data[currIndex], this.data[leftChild] = this.data[leftChild], this.data[currIndex]
				currIndex = leftChild
			} else {
				this.data[currIndex], this.data[rightChild] = this.data[rightChild], this.data[currIndex]
				currIndex = rightChild
			}
			continue
		}
		break
	}
	return
}

func topKFrequent(nums []int, k int) []int {
	result := []int{}
	frequency := make(map[int]int)
	for _, key := range nums {
		frequency[key]++
	}

	h := InitMaxHeap()

	for key, v := range frequency {
		node := &NumberFrequent{
			num:       key,
			frequency: v,
		}
		// fmt.Println(*node)
		h.insert(node)
	}

	// for _, v := range h.data {
	// 	fmt.Println(*v)
	// }

	i := 0
	for k > i {
		pop := h.pop()
		// for _, v := range h.data {
		// 	fmt.Println(*v)
		// }
		// fmt.Println("\n-----------")
		result = append(result, pop.num)
		i++
	}

	return result
}
