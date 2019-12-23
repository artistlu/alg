package main

import "fmt"

func main() {
	a := []int{8, 3, 9, 23, 98, 22, 11}
	r := countSort(a)
	fmt.Print(r)
	fmt.Println()
}

func countSort(a []int) []int {
	l := len(a)

	if l <= 1 {
		return a
	}

	mx := a[0]
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}

	c := make([]int, mx+1)
	for _, v := range a {
		c[v]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	r := make([]int, l)
	for _, v := range a {
		r[c[v]-1] = v
		c[v]--
	}

	return r
}
