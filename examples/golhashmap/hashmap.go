package main

import (
	"fmt"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
)

var (
	hashmap golhashmap.HashMap
	csv01   = "LedZep,Rock\nJohnnyCash,Country\nBach,Opera"
	json01  = "{ \"a\": \"b\", \"c\": \"d\" }"
)

func init() {
	hashmap = make(golhashmap.HashMap)
	hashmap["Bob"] = "Alice"
	hashmap["Eve"] = "Trudy"
}

func compare_map(map1 golhashmap.HashMap, map2 golhashmap.HashMap) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, val := range map1 {
		if val != map2[key] {
			return false
		}
	}
	return true
}

func TestCSVMap() {
	csv := golhashmap.HashMapToCSV(hashmap)
	expected_csv := "Bob,Alice\nEve,Trudy"

	hmap := golhashmap.CSVToHashMap(csv01)
	var expected_hmap golhashmap.HashMap
	expected_hmap = make(golhashmap.HashMap)
	expected_hmap["LedZep"] = "Rock"
	expected_hmap["JohnnyCash"] = "Country"
	expected_hmap["Bach"] = "Opera"

	if csv != expected_csv {
		panic("Conversion from HashMap to CSV is broken.")
	}
	if !compare_map(hmap, expected_hmap) {
		panic("Conversion from CSV to HashMap is broken.")
	}
}

func TestJSONMap() {
	json := golhashmap.HashMapToJSON(hashmap)
	expected_json := "{\"Bob\":\"Alice\",\"Eve\":\"Trudy\"}"

	hmap := golhashmap.JSONToHashMap(json01)
	var expected_hmap golhashmap.HashMap
	expected_hmap = make(golhashmap.HashMap)
	expected_hmap["a"] = "b"
	expected_hmap["c"] = "d"

	if json != expected_json {
		panic("Conversion from hmap to JSON is broken.")
	}
	if !compare_map(hmap, expected_hmap) {
		panic("Conversion from CSV to HashMap is broken.")
	}
}

func main() {
	TestCSVMap()
	TestJSONMap()
	fmt.Println("pass not panic")
}
