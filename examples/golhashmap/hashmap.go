package main

import (
	"fmt"

	golassert "github.com/abhishekkr/gol/golassert"
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

func TestHashMapKeys() {
	expected := []string{"Bob", "Eve"}
	golassert.AssertEqualStringArray(expected, hashmap.Keys())
}

func TestHashMapValues() {
	expected := []string{"Alice", "Trudy"}
	golassert.AssertEqualStringArray(expected, hashmap.Values())
}

func TestHashMapItems() {
	expected := [][]string{[]string{"Bob", "Alice"}, []string{"Eve", "Trudy"}}
	items := hashmap.Items()
	golassert.AssertEqualStringArray(expected[0], items[0])
	golassert.AssertEqualStringArray(expected[1], items[1])
	if len(expected) != len(items) {
		panic("Items length need to be same as expected.")
	}
}

func TestHashMapCount() {
	expected := len(hashmap)
	golassert.AssertEqual(expected, hashmap.Count())
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
	TestHashMapKeys()
	TestHashMapValues()
	TestHashMapItems()
	TestHashMapCount()
	fmt.Println("pass not panic")
}
