package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golcrypt"
)

func main() {
	txt := "whatever"

	txtAsByte := []byte(txt)
	key := golcrypt.MD5(txtAsByte)

	aesBlock := golcrypt.AESBlock{DataBlob: txtAsByte, Key: []byte(key), Cipher: nil}

	fmt.Println(key)

	aesBlock.Encrypt()
	aesBlock.DataBlob = nil
	fmt.Println(string(aesBlock.Cipher))

	aesBlock.Decrypt()
	aesBlock.Cipher = nil
	fmt.Println(string(aesBlock.DataBlob))
}
