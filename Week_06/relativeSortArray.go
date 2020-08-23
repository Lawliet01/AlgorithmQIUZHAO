package week6

func relativeSortArray(arr1 []int, arr2 []int) []int {
	n := 0
	for _, v := range arr2 {
		for i := n; i < len(arr1); i++ {
			if arr1[i] == v {
				arr1[i], arr1[n] = arr1[n], arr1[i]
				n++
			}
		}
	}

	quickSort(&arr1, n, len(arr1)-1)

	return arr1
}

func quickSort(arr *([]int), start int, pivot int) {
	if start >= pivot {
		return
	}
	counter := start
	for i := start; i < pivot; i++ {
		if (*arr)[i] < (*arr)[pivot] {
			(*arr)[counter], (*arr)[i] = (*arr)[i], (*arr)[counter]
			counter++
		}
	}

	(*arr)[counter], (*arr)[pivot] = (*arr)[pivot], (*arr)[counter]

	quickSort(arr, start, counter-1)
	quickSort(arr, counter+1, pivot)
}
