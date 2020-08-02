package week3

func subsets(nums []int) [][]int {
	res := [][]int{}
	addItem(0, &nums, []int{}, &res)
	return res
}

func addItem(i int, input *[]int, store []int, total *[][]int) {

	if i == len(*input) {
		*total = append(*total, store)
		return
	}

	copyStore := make([]int, len(store), len(store))
	copy(copyStore, store)

	addItem(i+1, input, append(store, (*input)[i]), total)
	addItem(i+1, input, copyStore, total)
}
