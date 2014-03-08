package golerror

import (
  "fmt"
)


/*
Prints provided error message and panics if rise value is True.
*/
func Boohoo(errstring string, rise bool){
  fmt.Println(errstring)
  if rise == true{ panic(errstring) }
}
