package main

import (
	"fmt"

	gollb "github.com/abhishekkr/gol/gollb"
)

func main() {
	testLoad()
	testLoadWithSeparator()
}

func testLoad() {
	services := map[string][]string{
		"/a": {
			"http://A1",
			"http://A2",
		},
		"/b": {
			"http://B1",
			"http://B2",
			"http://B3",
		},
		"/c": {
			"http://C1",
		},
	}

	rr := gollb.RoundRobin{}
	rr.Load(services)

	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("c: ", rr.GetBackend("/c"))
	fmt.Println("c: ", rr.GetBackend("/c"))
}

func testLoadWithSeparator() {
	services := map[string]string{
		"/a": "http://A1 http://A2",
		"/b": "http://B1 http://B2 http://B3",
		"/c": "http://C1",
	}

	rr := gollb.RoundRobin{}
	rr.LoadWithSeparator(services, " ")

	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("a: ", rr.GetBackend("/a"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("b: ", rr.GetBackend("/b"))
	fmt.Println("c: ", rr.GetBackend("/c"))
	fmt.Println("c: ", rr.GetBackend("/c"))
}
