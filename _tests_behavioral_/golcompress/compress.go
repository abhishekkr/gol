package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golcompress"
)

func main() {
	txt := "whatever----------------------------------------------------------"

	txtAsByte := []byte(txt)

	tiny, err := golcompress.Compress(txtAsByte, "brotli")
	fmt.Println(string(tiny))
	fmt.Println(len(tiny))
	fmt.Println(err)

	data, err := golcompress.Decompress(tiny, "brotli")
	fmt.Println(string(data))
	fmt.Println(len(data))
	fmt.Println(err)
}
