package main

import (
	"fmt"
	"strconv"

	golassert "github.com/abhishekkr/gol/golassert"
)

func Recover(msg string) {
	if r := recover(); r != nil {
		fmt.Printf("Passed PANIC for: %s With:\n%v\n\n", msg, r)
	} else {
		panic(fmt.Sprintf("Didn't panic where supposed to at %s", msg))
	}
}

func equalNil() {
	_, err := strconv.Atoi("1")
	golassert.AssertEqual(err, nil)

	defer Recover("equalNil")
	golassert.AssertEqual(1, nil)
}

func equalError() {
	_, err01 := strconv.Atoi("A")
	_, err02 := strconv.Atoi("A")
	golassert.AssertEqual(err01.Error(), err02.Error())

	defer Recover("equalError")
	golassert.AssertEqual(err01, nil)
}

func equalType() {
	golassert.AssertType(1, 1)

	defer Recover("equalType")
	golassert.AssertEqual(1, "1")
	golassert.AssertType(1, "1")
}

func equalString() {
	golassert.AssertEqual("1", "1")

	defer Recover("equalString")
	golassert.AssertEqual("1", "2")
}

func equalNumber() {
	golassert.AssertEqual(1, 1)

	defer Recover("equalNumber")
	golassert.AssertEqual(1, 2)
}

func equalStringArray() {
	var1 := []string{"a", "b"}
	var2 := []string{"b", "a"}
	var3 := []string{"b", "c"}
	var4 := []string{"a", "b"}
	golassert.AssertEqualStringArray(var1, var2)
	golassert.AssertEqualStringArray(var1, var4)

	defer Recover("equalStringArray")
	golassert.AssertEqualStringArray(var1, var3)
}

func main() {
	equalNil()
	equalError()
	equalType()
	equalString()
	equalNumber()
	equalStringArray()
}
