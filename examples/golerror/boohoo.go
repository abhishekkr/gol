package main

import (
	"fmt"

	golerror "../../golerror"
)

func main() {
	fmt.Println("Checking Boohoo without panic")
	golerror.Boohoo("I'm warning you.", false)
	fmt.Println("Checking Boohoo with panic")
	golerror.Boohoo("That's it.", true)
}
