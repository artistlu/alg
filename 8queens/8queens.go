package main

import "fmt"

func main() {
	// fmt.Print(queen)
	// fmt.Println()
	// return
	cal8queens(0)
}

var queen = [8]int{}

func cal8queens(row int) {

	if row == 8 {
		printQ(queen)
	}

	for c := 0; c < 8; c++ {
		if isOk(row, c) {
			queen[row] = c
			cal8queens(row + 1)
		}
	}
}

func isOk(row, c int) bool {
	leftUp := c - 1
	rightUp := c + 1
	for i := row - 1; i >= 0; i-- {
		if queen[i] == c {
			return false
		}
		if leftUp >= 0 && queen[i] == leftUp {
			return false
		}
		if rightUp < 8 && queen[i] == rightUp {
			return false
		}
		rightUp++
		leftUp--
	}

	return true
}

func printQ(q [8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == q[i] {
				fmt.Print("q ")
				continue
			}
			fmt.Print("* ")
		}

		fmt.Println()
		fmt.Println()
	}

	fmt.Print("--------------\n")
}
