package main

import (
	"errors"
	"fmt"
)

type BitMap struct {
	nbit  int
	bytes []rune
}

func (b *BitMap) Set(k int) (error, bool) {
	if k > b.nbit {
		return errors.New("out bitmap"), false
	}
	byteIndex := k / 32
	bitIndex := k % 32
	b.bytes[byteIndex] |= (1 << bitIndex)
	return nil, true
}

func (b *BitMap) Get(k int) (error, bool) {
	if k > b.nbit {
		return errors.New("out bitmap"), false
	}
	byteIndex := k / 32
	bitIndex := k % 32
	return nil, b.bytes[byteIndex]&(1<<bitIndex) != 0
}

func main() {
	word := make([]rune, 4)
	bm := BitMap{
		nbit:  100,
		bytes: word,
	}
	bm.Set(0)
	fmt.Print(bm.Get(32))
	fmt.Println()
	fmt.Print(bm.bytes)
}
