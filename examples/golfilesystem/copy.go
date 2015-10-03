package main

import (
	"flag"
	"fmt"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

var (
	fileSource      = flag.String("from", "", "Source")
	fileDestination = flag.String("to", "", "Destination")
)

func main() {
	flag.Parse()
	golfilesystem.CopyFile(*fileSource, *fileDestination)
	fmt.Println("Copied", *fileSource, "to", *fileDestination, ".")
}
