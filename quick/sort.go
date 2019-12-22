package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	quickSort(a)
	fmt.Print(a)
	fmt.Println()
}

func quickSort(a []int) {
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, s, e int) {
	if s >= e {
		return
	}
	p := partition(a, s, e)
	fmt.Printf("[%d-%d-%d]\n", s, p, e)
	qsort(a, s, p-1)
	qsort(a, p+1, e)
}

func partition(a []int, s, e int) int {
	i, j := s, s
	for j <= e-1 {
		if a[j] < a[e] {
			a[i], a[j] = a[j], a[i]
			i++
		}
		j++
	}
	a[i], a[e] = a[e], a[i]
	return i
}
