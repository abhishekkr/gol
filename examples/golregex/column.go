package main

import (
	"fmt"

	golassert "github.com/abhishekkr/gol/golassert"

	golregex "github.com/abhishekkr/gol/golregex"
)

func TestColumn() {
	fmt.Println("testing: golregex.Column")
	golassert.AssertEqual(golregex.Column("tsds", "-", 1), "tsds")
	golassert.AssertEqual(golregex.Column("tsds-csv", "-", 1), "tsds")
	golassert.AssertEqual(golregex.Column("tsds-csv-now", "-", 1), "tsds")
	golassert.AssertEqual(golregex.Column("tsds", "-", 2), "")
	golassert.AssertEqual(golregex.Column("tsds-csv", "-", 2), "csv")
	golassert.AssertEqual(golregex.Column("tsds-csv-now", "-", 2), "csv")
	golassert.AssertEqual(golregex.Column("tsds", "-", 3), "")
	golassert.AssertEqual(golregex.Column("tsds-csv", "-", 3), "")
	golassert.AssertEqual(golregex.Column("tsds-csv-now", "-", 3), "now")
}

func main() {
	TestColumn()
}
