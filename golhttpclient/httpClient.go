package golhttpclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	SkipSSLVerify bool
)

type HTTPRequest struct {
	Method      string
	Url         string
	GetParams   map[string]string
	HTTPHeaders map[string]string
	Body        *bytes.Buffer
}

func customRoundTripper() (customTransport http.RoundTripper) {
	customTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: SkipSSLVerify},
		ExpectContinueTimeout: 1 * time.Second,
	}
	/*
		var ProxyFromEnvironment *http.Request
		if ProxyFromEnvironment != nil {
			*customTransport.Proxy = ProxyFromEnvironment
		}
	*/
	return
}

func UrlRedirectTo(url string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Transport: customRoundTripper(),
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Println(err)
	}

	return resp.Request.URL.String()
}

func LinkExists(url string) bool {
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: customRoundTripper(),
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
	requestUrl := fmt.Sprintf("%s?%s", httpRequest.Url, getParamsURI)
	url, err := url.Parse(requestUrl)

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
	for header, value := range httpRequest.HTTPHeaders {
		req.Header.Add(header, value)
	}
	return
}

func (httpRequest *HTTPRequest) httpResponse() (resp *http.Response, err error) {
	httpClient := &http.Client{
		Transport: customRoundTripper(),
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

func (httpRequest *HTTPRequest) Response() (*http.Response, error) {
	return httpRequest.httpResponse()
}

func (httpRequest *HTTPRequest) httpResponseBytes() (body []byte, err error) {
	resp, err := httpRequest.httpResponse()
	if err != nil {
		log.Println(err)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (httpRequest *HTTPRequest) httpResponseBody() (body string, err error) {
	bodyText, err := httpRequest.httpResponseBytes()
	if err == nil {
		body = string(bodyText)
	} else {
		log.Println(err)
		return
	}
	return
}

func (httpRequest *HTTPRequest) Http(httpMethod string) (resp *http.Response, err error) {
	httpRequest.Method = httpMethod
	resp, err = httpRequest.httpResponse()
	return
}

func (httpRequest *HTTPRequest) GetBytes() (body []byte, err error) {
	httpRequest.Method = "GET"
	body, err = httpRequest.httpResponseBytes()
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
