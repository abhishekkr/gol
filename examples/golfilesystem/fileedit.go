package main

import (
	"flag"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

var (
	filename      = flag.String("file", "", "create/append to file")
)

func main() {
	flag.Parse()
	golfilesystem.AppendToFile(*filename, "what")
	golfilesystem.AppendToFile(*filename, "why")
}
