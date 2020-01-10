package golnet

import (
	"net/http"
	"time"
)

func IsHttpAvailable(url string, timeoutInSecond int) bool {
	timeout := time.Duration(timeoutInSecond) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Head(url)
	if err != nil {
		return false
	}
	return true
}
