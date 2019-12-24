package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := bSearch(a, 10)
	fmt.Println(i)
	fmt.Println()
}

func bSearch(a []int, b int) int {
	l := 0
	h := len(a) - 1

	for l <= h {
		m := l + ((h - l) >> 2)
		fmt.Printf("%d\n", m)
		if a[m] == b {
			return m
		} else if a[m] > b {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
