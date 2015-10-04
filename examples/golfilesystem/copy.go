package main

import (
	"flag"
	"fmt"

	//golfilesystem "github.com/abhishekkr/gol/golfilesystem"
	golfilesystem "../../golfilesystem"
)

var (
	fileSource      = flag.String("from", "", "Source")
	fileDestination = flag.String("to", "", "Destination")
)

func main() {
	flag.Parse()
	if err := golfilesystem.CopyFile(*fileSource, *fileDestination); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("Copied", *fileSource, "to", *fileDestination, ".")
	}
}
