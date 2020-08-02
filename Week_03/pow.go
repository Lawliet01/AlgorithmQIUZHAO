package week3

func myPow(x float64, n int) float64 {
	val := x
	if n == 0 {
		return 1
	}
	if n < 0 {
		val = 1 / val
		n = -n
	}

	store := make(map[int]float64)
	store[0] = val
	store[1] = val * val

	return helper(1, n, val, store)
}

func helper(startN int, endN int, value float64, store map[int]float64) float64 {

	diff := endN - startN
	v, ok := store[diff]

	if ok {
		return v
	}

	centerPoint := (endN + startN) / 2

	res := helper(startN, centerPoint, value, store) * helper(centerPoint+1, endN, value, store)

	store[diff] = res

	return res
}
