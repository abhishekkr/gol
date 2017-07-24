package golhttpclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HTTPRequest struct {
	Method      string
	Url         string
	GetParams   map[string]string
	HTTPHeaders map[string]string
	Body        *bytes.Buffer
}

func UrlRedirectTo(url string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Println(err)
	}

	return resp.Request.URL.String()
}

func LinkExists(url string) bool {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil || response.StatusCode > 399 {
		return false
	}
	return true
}

func (httpRequest *HTTPRequest) getURL() (url *url.URL) {
	var getParamsURI string
	for key, val := range httpRequest.GetParams {
		if getParamsURI == "" {
			getParamsURI = fmt.Sprintf("%s=%s", key, val)
		} else {
			getParamsURI = fmt.Sprintf("%s&%s=%s", getParamsURI, key, val)
		}
	}
	request_url := fmt.Sprintf("%s?%s", httpRequest.Url, getParamsURI)
	url, err := url.Parse(request_url)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (httpRequest *HTTPRequest) setHttpHeaders(req *http.Request) (err error) {
	basicAuth := strings.Split(httpRequest.HTTPHeaders["basicAuth"], ":")
	if len(basicAuth) > 1 {
		apiUsername, apiPassword := basicAuth[0], strings.Join(basicAuth[1:], ":")
		req.SetBasicAuth(apiUsername, apiPassword)
	}
	return
}

func (httpRequest *HTTPRequest) httpResponse() (resp *http.Response, err error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	var req *http.Request
	if httpRequest.Body == nil {
		req, err = http.NewRequest(httpRequest.Method, "", nil)
	} else {
		req, err = http.NewRequest(httpRequest.Method, "", httpRequest.Body)
	}
	if err != nil {
		return
	}
	req.URL = httpRequest.getURL()
	httpRequest.setHttpHeaders(req)

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (httpRequest *HTTPRequest) httpResponseBody() (body string, err error) {
	resp, err := httpRequest.httpResponse()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		body = string(bodyText)
	} else {
		log.Println(err)
	}
	return
}

func (httpRequest *HTTPRequest) Http(httpMethod string) (resp *http.Response, err error) {
	httpRequest.Method = httpMethod
	resp, err = httpRequest.httpResponse()
	return
}

func (httpRequest *HTTPRequest) Get() (body string, err error) {
	httpRequest.Method = "GET"
	body, err = httpRequest.httpResponseBody()
	return
}

func (httpRequest *HTTPRequest) Put() (body string, err error) {
	httpRequest.Method = "PUT"
	body, err = httpRequest.httpResponseBody()
	return
}

func (httpRequest *HTTPRequest) Post() (body string, err error) {
	httpRequest.Method = "POST"
	body, err = httpRequest.httpResponseBody()
	return
}

func (httpRequest *HTTPRequest) Delete() (body string, err error) {
	httpRequest.Method = "DELETE"
	body, err = httpRequest.httpResponseBody()
	return
}
