package main

import (
	"flag"
	"fmt"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

var (
	fileSource = flag.String("read", "", "Source")
)

func main() {
	flag.Parse()
	buffer, err := golfilesystem.FileBuffer(*fileSource)
	if err != nil {
		fmt.Println(">>", *fileSource)
		fmt.Println("ERROR:", err)
	}
	fmt.Println("Bytes Length:", buffer.Len())
	fmt.Println(buffer.String())
}
