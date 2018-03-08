package main

import (
	"log"

	"github.com/abhishekkr/gol/golhttpclient"
)

func main() {
	req := golhttpclient.HTTPRequest{}
	req.Url = "https://www.google.com"
	golhttpclient.SkipSSLVerify = true
	response, err := req.Get()
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}
