package main

import (
	"fmt"

	"github.com/abhishekkr/gol/golcrypt"
)

func main() {
	txt := "whatever"

	txtAsByte := []byte(txt)
	key := []byte("secret-key-need-to-be-bigger")

	cipher, err := golcrypt.Encrypt(txtAsByte, key, "aes")
	fmt.Println(string(cipher))
	fmt.Println(err)

	data, err := golcrypt.Decrypt(cipher, key, "aes")
	fmt.Println(string(data))
	fmt.Println(err)
}
