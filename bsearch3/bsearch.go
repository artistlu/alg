package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 6, 8, 8, 8, 9, 10}
	i := bsearch(a, 8)
	fmt.Println(i)
	fmt.Println()
}

func bsearch(a []int, s int) int {

	l := 0
	h := len(a) - 1

	for l <= h {
		m := l + ((h - l) >> 1)
		if a[m] > s {
			h = m - 1
		} else if a[m] < s {
			l = m + 1
		} else {
			if m == len(a)-1 || a[m+1] != s {
				return m
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
