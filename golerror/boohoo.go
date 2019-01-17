package golerror

import (
	"fmt"
	"os"
)

var (
	BoohooType = os.Getenv("BOOHOO")
)

/*
Boohoo prints provided error message and panics if rise value is True.
*/
func Boohoo(errstring string, rise bool) {
	if BoohooType == "stdout" {
		fmt.Println(errstring)
	}
	if rise {
		panic(errstring)
	}
}
