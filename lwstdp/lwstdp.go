package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Print("hello\n")

	ed := lwstdp("abf", "ab")
	fmt.Println(ed)
}

func lwstdp(a, b string) int {
	table := make([][]int, 1)
	firstRow := make([]int, 1)
	for i := range a {
		var t int
		if a[i] == b[0] {
			t = i
		} else if i != 0 {
			t = firstRow[i-1] + 1
		} else {
			t = 1
		}

		if i == 0 {
			firstRow[0] = t
		} else {
			firstRow = append(firstRow, t)
		}

	}
	table[0] = firstRow

	for j := range b {
		if j == 0 {
			continue
		}

		next := make([]int, 1)
		if a[0] == b[j] {
			next[0] = j
		} else {
			next[0] = table[j-1][0] + 1
		}

		table = append(table, next)
	}

	for i := 1; i < len(b); i++ {
		for j := 1; j < len(a); j++ {
			if a[j] == b[i] {
				x := table[i-1][j] + 1
				y := table[i][j-1] + 1
				z := table[i-1][j-1]
				table[i] = append(table[i], min(x, y, z))
			} else {
				e := table[i-1][j] + 1
				f := table[i][j-1] + 1
				g := table[i-1][j-1] + 1
				table[i] = append(table[i], min(e, f, g))
			}
		}
	}

	fmt.Println(table)
	return table[len(b)-1][len(a)-1]
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}

	if x > z {
		x = z
	}

	return x
}
