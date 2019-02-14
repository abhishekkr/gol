package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golconv"
)

func testStringToUint64() {
	u := golconv.StringToUint64("100", 10)
	if u != uint64(100) {
		panic("StringToUint64 for correct string failed")
	}

	u = golconv.StringToUint64("00", 10)
	if u != uint64(0) {
		panic("StringToUint64 for correct string failed")
	}

	u = golconv.StringToUint64("i00", 10)
	if u != uint64(10) {
		panic("StringToUint64 for wrong string failed")
	}
}

func testStringToInt() {
	i := golconv.StringToInt("100", 10)
	if i != 100 {
		panic("StringToInt for correct string failed")
	}

	i = golconv.StringToInt("00", 10)
	if i != 0 {
		panic("StringToInt for correct string failed")
	}

	i = golconv.StringToInt("i00", 10)
	if i != 10 {
		panic("StringToInt for wrong string failed")
	}
}

func main() {
	testStringToUint64()
	testStringToInt()
	fmt.Println("+ golconv passed")
}
