package golhttpclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getURL(baseURL string, getParams map[string]string) (url *url.URL) {
	var getParamsURI string
	for key, val := range getParams {
		if getParamsURI == "" {
			getParamsURI = fmt.Sprintf("%s=%s", key, val)
		} else {
			getParamsURI = fmt.Sprintf("%s&%s=%s", getParamsURI, key, val)
		}
	}
	request_url := fmt.Sprintf("%s?%s", baseURL, getParamsURI)
	url, err := url.Parse(request_url)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func setHttpHeaders(req *http.Request, httpHeaders map[string]string) (err error) {
	basicAuth := strings.Split(httpHeaders["basicAuth"], ":")
	if len(basicAuth) > 1 {
		apiUsername, apiPassword := basicAuth[0], strings.Join(basicAuth[1:], ":")
		req.SetBasicAuth(apiUsername, apiPassword)
	}
	return
}

func httpClient(httpMethod string, baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
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

	req, err := http.NewRequest(httpMethod, "", nil)
	if err != nil {
		return
	}
	req.URL = getURL(baseURL, getParams)
	setHttpHeaders(req, httpHeaders)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		body = string(bodyText)
	} else {
		log.Println(err)
	}

	return
}

func HttpGet(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	body, err = httpClient("GET", baseURL, getParams, httpHeaders)
	return
}

func HttpPut(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	// need to handle PUT body content
	body, err = httpClient("PUT", baseURL, getParams, httpHeaders)
	return
}

func HttpPost(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	// need to handle POST body content
	body, err = httpClient("POST", baseURL, getParams, httpHeaders)
	return
}

func HttpDelete(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	body, err = httpClient("DELETE", baseURL, getParams, httpHeaders)
	return
}
