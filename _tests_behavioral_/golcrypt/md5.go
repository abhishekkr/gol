package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golcrypt"
)

func main() {
	txt := "whatever"
	txtAsByte := []byte(txt)
	md5sum := golcrypt.MD5(txtAsByte)

	fmt.Println(md5sum)
}
