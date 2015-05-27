package main

import (
	"fmt"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

func main() {
	exists := "."
	exists_not := "..."
	if !golfilesystem.PathExists(exists) {
		panic(fmt.Sprintf("%s does exists!", exists))
	}
	if golfilesystem.PathExists(exists_not) {
		panic(fmt.Sprintf("%s doesn't exist!", exists_not))
	}
	fmt.Println("No need to panic.")
}
