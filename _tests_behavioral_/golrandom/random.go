package main

import (
	"fmt"

	golrandom "github.com/abhishekkr/gol/golrandom"
)

func main() {
	fmt.Println(golrandom.Name(16))
	fmt.Println(golrandom.Token(16))
}
