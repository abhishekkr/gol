package main

import (
	"log"

	"github.com/abhishekkr/gol/golhttpclient"
)

func main() {
	req := golhttpclient.HTTPRequest{}
	req.Url = "https://status.cloud.google.com/incidents.json"
	golhttpclient.SkipSSLVerify = true
	response, err := req.Get()
	if err != nil {
		log.Println(err)
	}

	responseB, err := req.GetBytes()
	if err != nil {
		log.Println(err)
	}

	if response != string(responseB) || len(response) == 0 {
		log.Fatalln("different responses")
	}

	log.Println("pass.")
}
