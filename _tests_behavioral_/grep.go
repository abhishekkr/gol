package main

import (
	"flag"
	"fmt"

	golbin "github.com/abhishekkr/gol/golbin"
)

func main() {
	flag.Parse()

	filepath := flag.Arg(0)
	fmt.Println(filepath)
	lines, file_err := golbin.Cat(filepath)
	if file_err != nil {
		fmt.Println(file_err)
	}

	match, _ := golbin.Grep("go", lines)
	fmt.Println(match)
}
