package main

import (
	"fmt"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

func main() {
	exists := "."
	existsNot := "..."
	if !golfilesystem.PathExists(exists) {
		panic(fmt.Sprintf("%s does exists!", exists))
	}
	if golfilesystem.PathExists(existsNot) {
		panic(fmt.Sprintf("%s doesn't exist!", existsNot))
	}
	fmt.Println("No need to panic.")
}
