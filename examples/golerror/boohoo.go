package main

import (
	"fmt"

	golerror "github.com/abhishekkr/gol/golerror"
)

func main() {
	fmt.Println("Checking Boohoo without panic")
	golerror.Boohoo("I'm warning you.", false)
	fmt.Println("Checking Boohoo with panic")
	golerror.Boohoo("That's it.", true)
}
