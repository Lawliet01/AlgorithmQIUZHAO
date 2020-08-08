package week4

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1) + 1
	n := len(text2) + 1

	memo := make([]int, m*n, m*n)

	for i, v := range text1 {
		for j, v2 := range text2 {
			if v == v2 {
				memo[(i+1)*n+(j+1)] = memo[i*n+j] + 1
				continue
			}

			store1 := memo[i*n+(j+1)]
			store2 := memo[(i+1)*n+j]

			if store1 > store2 {
				memo[(i+1)*n+(j+1)] = store1
			} else {
				memo[(i+1)*n+(j+1)] = store2
			}
		}
	}

	return memo[m*n-1]
}
