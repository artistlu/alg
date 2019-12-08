package main

import "fmt"

var maxW int

func main() {
	n := 10
	a := make([]int, n)
	w := 101
	a = []int{102, 9, 13, 15, 19, 30, 20, 8, 6, 99}

	f(0, 0, a, n, w)

	fmt.Print(maxW)
	fmt.Println()
}

func f(i int, cw int, a []int, n int, w int) {
	if i == n || cw == w {
		if maxW < cw {
			maxW = cw
		}
		return
	}

	f(i+1, cw, a, n, w)
	if cw+a[i] <= w {
		f(i+1, cw+a[i], a, n, w)
	}
}
