package main

import "fmt"

func main() {
	a := []int{3, 2, 1, 7, 6, 2, 9}
	fmt.Print(a)
	fmt.Println()
	insertSort(a)
	// fmt.Print(a)
	// fmt.Println()
	// fmt.Print(b)
	fmt.Println()
}

func insertSort(a []int) []int {

	l := len(a)

	for j := 1; j < l; j++ {
		v := a[j]
		i := j - 1
		for ; i >= 0; i-- {
			if a[i] > v {
				a[i+1] = a[i]
			} else {
				break
			}
		}
		a[i+1] = v
		// fmt.Print(a)
		// fmt.Println()

	}

	return a
}
