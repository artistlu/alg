package main

import (
	"fmt"
)

func main() {
	// bm := bitmap.NewBitMap(100)
	// // fmt.Printf("%b\n", 0x08)
	// fmt.Print(bm)
	a := 6

	fmt.Printf(" %b\n", 255)
	fmt.Printf(" %b\n", 1<<a+^(1<<a))
	fmt.Printf(" %b\n", 255 & ^(1<<a))
	// fmt.Printf("%b\n", a)
	// fmt.Printf("%b\n", a&7)

}
