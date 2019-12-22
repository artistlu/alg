package main

import "fmt"

func main() {
	a := []int{4, 5, 9, 3, 2, 1, 0, 8, 4}
	mergeSort(a)
	fmt.Print(a)
	fmt.Println()
}

func mergeSort(a []int) []int {
	msort(a, 0, len(a)-1)
	return a
}

func msort(a []int, s, e int) {
	if s >= e {
		return
	}

	m := s + (e-s)/2
	// fmt.Print(m)
	// fmt.Println()
	msort(a, s, m)
	msort(a, m+1, e)
	mmerge(a, s, m, e)
}

func mmerge(a []int, s, m, e int) {
	fmt.Printf("merge [%d-%d] [%d-%d]\n", s, m, m+1, e)
	t := make([]int, e-s+1)

	i := s
	j := m + 1
	k := 0
	for i <= m && j <= e {
		if a[i] <= a[j] {
			t[k] = a[i]
			i++
		} else {
			t[k] = a[j]
			j++
		}
		k++
	}

	ls := i
	le := m

	if i > m {
		ls = j
		le = e
	}
	for i := ls; i <= le; i++ {
		t[k] = a[i]
		k++
	}
	k = 0
	for i := s; i <= e; i++ {
		a[i] = t[k]
		k++
	}
}
