package week2

type MinHeap struct {
	data []int
}

func InitHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (this *MinHeap) insert(n int) {
	this.data = append(this.data, n)
	currIndex := len(this.data) - 1
	for {
		parentNodeIndex := (currIndex - 1) / 2

		if this.data[parentNodeIndex] < this.data[currIndex] || currIndex == 0 {
			return
		}

		this.data[parentNodeIndex], this.data[currIndex] = this.data[currIndex], this.data[parentNodeIndex]
		currIndex = parentNodeIndex
	}
}

func (this *MinHeap) pop() (result int) {
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
			if this.data[leftChild] < this.data[currIndex] {
				this.data[currIndex], this.data[leftChild] = this.data[leftChild], this.data[currIndex]
			}
			break
		}

		if this.data[currIndex] > this.data[leftChild] || this.data[currIndex] > this.data[rightChild] {
			if this.data[leftChild] < this.data[rightChild] {
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

func (this *MinHeap) findMin(numOfMin int) (result []int) {
	for numOfMin > len(result) {
		result = append(result, this.pop())
	}
	return
}

func getLeastNumbers(arr []int, k int) []int {
	h := InitHeap()
	for _, v := range arr {
		h.insert(v)
	}
	return h.findMin(k)
}
