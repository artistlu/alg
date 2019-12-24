package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 6, 8, 8, 8, 9, 10}
	i := bsearch(a, 100)
	fmt.Println(i)
	fmt.Println()
}

//找下标最小的值
func bsearch(a []int, b int) int {
	l := 0
	h := len(a) - 1
	for l <= h {
		m := l + ((h - l) >> 1)

		if a[m] > b {
			h = m - 1
		} else if a[m] < b {
			l = m + 1
		} else {
			if m == 0 || a[m-1] != b {
				return m
			} else {
				h = m - 1
			}
		}
	}

	return -1
}
