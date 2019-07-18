package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golerror"
)

func main() {
	e := golerror.Error(127, "say what")
	fmt.Println(e)
}
