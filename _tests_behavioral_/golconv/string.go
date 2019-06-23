package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golassert"
	"github.com/abhishekkr/gol/golconv"
)

func testStringToUint64() {
	u := golconv.StringToUint64("100", 10)
	golassert.AssertEqual(u, uint64(100))

	u = golconv.StringToUint64("00", 10)
	golassert.AssertEqual(u, uint64(0))

	u = golconv.StringToUint64("i00", 10)
	golassert.AssertEqual(u, uint64(10))
}

func testStringToInt() {
	i := golconv.StringToInt("100", 10)
	golassert.AssertEqual(i, 100)

	i = golconv.StringToInt("00", 10)
	golassert.AssertEqual(i, 0)

	i = golconv.StringToInt("i00", 10)
	golassert.AssertEqual(i, 10)
}

func testStringToBool() {
	b := golconv.StringToBool("true", false)
	golassert.AssertEqual(b, true)

	b = golconv.StringToBool("false", true)
	golassert.AssertEqual(b, false)

	b = golconv.StringToBool("whatever", true)
	golassert.AssertEqual(b, true)

	b = golconv.StringToBool("whatever", false)
	golassert.AssertEqual(b, false)
}

func main() {
	testStringToUint64()
	testStringToInt()
	testStringToBool()
	fmt.Println("+ golconv passed")
}
