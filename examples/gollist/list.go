package main

import (
	"fmt"

	"github.com/abhishekkr/gol/gollist"
)

var (
	someList = []string{"Bob", "Alice", "Eve", "Trudy", "Jack"}
	csv01    = "abba,dabba,chabba"
	csv02    = "LedZep,Rock\nJohnnyCash,Country\nBach,Opera"
	csv03    = "1,10,100,1000,10000"
	json01   = "[ \"I\", \"U\", \"We\" ]"
	json02   = "{ \"a\": \"b\", \"c\": \"d\" }"
)

func compare_list(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}
	for idx, val := range list1 {
		if val != list2[idx] {
			return false
		}
	}
	return true
}

func compare_numbers(list1 []int, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}
	for idx, val := range list1 {
		if val != list2[idx] {
			return false
		}
	}
	return true
}

func TestCSVMap() {
	csv := gollist.ListToCSV(someList)
	expected_csv := "Bob,Alice,Eve,Trudy,Jack"

	list01 := gollist.CSVToList(csv01)
	expected_list01 := []string{"abba", "dabba", "chabba"}

	list02 := gollist.CSVToList(csv02)
	expected_list02 := []string{"LedZep", "Rock", "JohnnyCash", "Country", "Bach", "Opera"}

	list03, _ := gollist.CSVToNumbers(csv03)
	expected_list03 := []int{1, 10, 100, 1000, 10000}

	if csv != expected_csv {
		panic("Conversion from List to CSV is broken.")
	}
	if !(compare_list(list01, expected_list01) && compare_list(list02, expected_list02)) {
		panic("Conversion from CSV to List is broken.")
	}
	if !compare_numbers(list03, expected_list03) {
		panic("Conversion from CSV to List is broken.")
	}
}

func TestJSONMap() {
	json := gollist.ListToJSON(someList)
	expected_json := "[\"Bob\",\"Alice\",\"Eve\",\"Trudy\",\"Jack\"]"

	list01 := gollist.JSONToList(json01)
	expected_list01 := []string{"I", "U", "We"}

	list02 := gollist.JSONToList(json02)
	expected_list02 := []string{}

	if json != expected_json {
		panic("Conversion from list to JSON is broken.")
	}
	if !(compare_list(list01, expected_list01) && compare_list(list02, expected_list02)) {
		panic("Conversion from CSV to List is broken.")
	}
}

func main() {
	TestCSVMap()
	TestJSONMap()
	fmt.Println("pass not panic")
}
