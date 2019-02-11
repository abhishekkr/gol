package main

import "github.com/abhishekkr/gol/golconv"

func main() {
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
