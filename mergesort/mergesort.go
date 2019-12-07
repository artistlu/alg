package main

import "fmt"

var re int

func main() {
	s := []int{7, 5, 2, 34, 66, 22, 66, 1, 3, 8, 0, 8}
	mergesort(&s, 0, len(s)-1)
	fmt.Print(s)
	fmt.Println()
	fmt.Print(re)
	fmt.Println()
}

func mergesort(s *[]int, p int, q int) {
	if p >= q {
		return
	}

	m := p + (q-p)/2
	mergesort(s, p, m)
	mergesort(s, m+1, q)
	tmerge(s, p, m, q)
}

func tmerge(s *[]int, p int, m int, q int) {
	i, j := p, m+1
	tmp := make([]int, 0)
	for i <= m && j <= q {
		if (*s)[i] <= (*s)[j] {
			tmp = append(tmp, (*s)[i])
			i++
		} else {
			re += m - i + 1
			tmp = append(tmp, (*s)[j])
			j++
		}
	}

	if i <= m {
		tmp = append(tmp, (*s)[i:m+1]...)
	}

	if j <= q {
		tmp = append(tmp, (*s)[j:q+1]...)
	}

	// fmt.Print((*s)[p : m+1])
	// fmt.Println()
	// fmt.Print((*s)[m+1 : q+1])
	// fmt.Println()
	// fmt.Print("tmp:", tmp)
	// fmt.Println()

	for k, v := range tmp {
		(*s)[p+k] = v
	}
}
