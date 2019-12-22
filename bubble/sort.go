package main

import "fmt"

func main() {
	a := []int{0, 0, 0}
	b := bubbleSort(a)
	fmt.Print(a)
	fmt.Println()
	fmt.Print(b)
	fmt.Println()
}

func bubbleSort(a []int) []int {

	l := len(a)

	for j := l - 1; j > 0; j-- {
		c := true
		for i := 0; i < j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				c = false
			}
		}
		if c {
			break
		}

	}

	return a
}
