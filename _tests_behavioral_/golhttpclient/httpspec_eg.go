package main

import (
	"fmt"
	"regexp"

	"github.com/abhishekkr/gol/golassert"
	"github.com/abhishekkr/gol/golhttpclient"
)

var sampleX = `
POST /what/is/path HTTP/1.1
Host: 10.0.0.1:8080
Content-Type: application/json
cache-control: no-cache
Postman-Token: xxx

{
    "start_date": "2029-04-19T00:00:00Z",
    "end_date": "2039-04-19T00:00:00Z",
    "name": "Alice",
    "friends": [
        "Bob",
        "Eve"
    ]
}
`

var sampleY = `GET /?statuses=read,liked&amp; page=1 HTTP/1.0
Host: 1.1.2.0
cache-control: no-cache

`

func main() {
	checkSpecX()
	checkSpecY()
}

func checkSpecX() {
	spec := &golhttpclient.HTTPSpec{Spec: sampleX}
	spec.Parse()

	golassert.Equal(spec.Client.Method, "POST")
	golassert.Equal(spec.Client.Protocol, "HTTP/1.1")
	golassert.Equal(spec.Client.Url, "/what/is/path")
	golassert.EqualStringMap(spec.Client.GetParams, map[string]string{})
	golassert.EqualStringMap(spec.Client.HTTPHeaders,
		map[string]string{"Content-Type": "application/json", "Host": "10.0.0.1:8080", "Postman-Token": "xxx", "cache-control": "no-cache"},
	)
	match, _ := regexp.MatchString("\\{.*start_date.*end_date.*name.*friends.*\\}", spec.Client.Body.String())
	golassert.Equal(match, true)
	fmt.Println("passed for sampleX")
}

func checkSpecY() {
	spec := &golhttpclient.HTTPSpec{Spec: sampleY}
	spec.Parse()

	golassert.Equal(spec.Client.Method, "GET")
	golassert.Equal(spec.Client.Protocol, "HTTP/1.0")
	golassert.Equal(spec.Client.Url, "/")
	golassert.EqualStringMap(spec.Client.GetParams,
		map[string]string{"Host": "1.1.2.0", "cache-control": "no-cache"},
	)
	golassert.EqualStringMap(spec.Client.HTTPHeaders,
		map[string]string{"Host": "1.1.2.0", "cache-control": "no-cache"},
	)
	golassert.Equal("", spec.Client.Body.String())
	fmt.Println("passed for sampleY")
}
