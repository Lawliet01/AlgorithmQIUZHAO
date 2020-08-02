package week3

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	for i := 0; i < len(s) && res < len(g); i++ {
		if s[i] >= g[res] {
			res++
		}
	}
	return res
}
