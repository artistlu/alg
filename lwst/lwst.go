package main

import "fmt"

var a string = "zhbnglu"
var b string = "zhangjun"
var alen int = len(a) - 1
var blen int = len(b) - 1
var minEdist int

func main() {
	fmt.Print(minEdist)
	fmt.Println()
	lwst(0, 0, 0)
	fmt.Print(minEdist)
	fmt.Println()
}

func lwst(i, j, edist int) {
	if i == alen || j == blen {
		if i < alen {
			edist += alen - i
		}

		if j < blen {
			edist += blen - j
		}

		if minEdist == 0 {
			minEdist = edist
		} else if edist < minEdist {
			minEdist = edist
		}
		return

	}
	if a[i] == b[j] {
		lwst(i+1, j+1, edist)
	} else {
		lwst(i+1, j, edist+1)
		lwst(i, j+1, edist+1)
		lwst(i+1, j+1, edist+1)
	}
}
