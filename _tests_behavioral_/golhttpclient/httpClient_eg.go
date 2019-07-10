package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/abhishekkr/gol/golhttpclient"
	handler "github.com/abhishekkr/http200/handler"
)

var (
	server *http.Server

	listenAt = "127.0.0.1:65500"
	quit     = make(chan bool)
)

func checkForHttpMethod(req golhttpclient.HTTPRequest, method string) {
	req.Method = method
	resp, err := req.Response()
	if err != nil || resp.StatusCode != 200 {
		log.Fatalf("%s failed\n", method)
	}
}

func dummyServer() {
	server := &http.Server{
		Addr:    listenAt,
		Handler: handler.AppHandler(),
	}
	server.SetKeepAlivesEnabled(false)

	done := make(chan bool)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if errX := server.Shutdown(ctx); errX != nil {
			panic(errX)
		}
		close(done)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	<-done
	log.Println("dummy server closed")
}

func waitForServer() {
	timeout := 10 * time.Second
	for {
		conn, _ := net.DialTimeout("tcp", listenAt, timeout)
		if conn != nil {
			fmt.Println("dummy server started")
			conn.Close()
			break
		}
	}
}

func main() {
	go dummyServer()
	waitForServer()

	req := golhttpclient.HTTPRequest{}
	req.Url = fmt.Sprintf("http://%s", listenAt)
	golhttpclient.SkipSSLVerify = true
	response, err := req.Get()
	if err != nil {
		log.Println(err)
	}

	responseB, err := req.GetBytes()
	if err != nil {
		log.Println(err)
	}

	responseC, err := req.Response()
	if err != nil {
		log.Println(err)
	}
	responseCBody, err := ioutil.ReadAll(responseC.Body)
	if err != nil {
		log.Println(err)
	}

	if response != string(responseB) || len(response) == 0 {
		log.Fatalln("different responses")
	}

	if response != string(responseCBody) || len(response) == 0 {
		log.Fatalln("different responses")
	}

	req.Url += "/200"
	checkForHttpMethod(req, "GET")
	checkForHttpMethod(req, "POST")
	checkForHttpMethod(req, "PUT")
	checkForHttpMethod(req, "PATCH")
	checkForHttpMethod(req, "DELETE")
	checkForHttpMethod(req, "HEAD")
	checkForHttpMethod(req, "OPTIONS")

	quit <- true
	log.Println("pass.")
}
