package main

import (
	"fmt"
	"regexp"

  golregex "github.com/abhishekkr/gol/golregex"
)

func main() {
	fmt.Println(golregex.Column("tsds", "-", 1))
	fmt.Println(golregex.Column("tsds-csv", "-", 1))
	fmt.Println(golregex.Column("tsds-csv-now", "-", 1))
	fmt.Println(golregex.Column("tsds", "-", 2))
	fmt.Println(golregex.Column("tsds-csv", "-", 2))
	fmt.Println(golregex.Column("tsds-csv-now", "-", 2))
	fmt.Println(golregex.Column("tsds", "-", 3))
	fmt.Println(golregex.Column("tsds-csv", "-", 3))
	fmt.Println(golregex.Column("tsds-csv-now", "-", 3))
}
