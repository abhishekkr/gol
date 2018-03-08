package main

import (
	"fmt"
	"time"

	"github.com/abhishekkr/gol"
)

func waitNShout() {
	fmt.Println("Gol")
	time.Sleep(1 * time.Second)
}

func main() {
	gol.Gol(waitNShout)
}
