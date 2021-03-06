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
	Protocol    string
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

func UrlRedirectTo(uri string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Transport: customRoundTripper(),
	}

	resp, err := client.Get(uri)

	if err != nil {
		log.Println(err)
	}

	return resp.Request.URL.String()
}

func LinkExists(uri string) bool {
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: customRoundTripper(),
	}
	response, err := netClient.Get(uri)
	if err != nil || response.StatusCode > 399 {
		return false
	}
	return true
}

func (httpRequest *HTTPRequest) getURL() (uri *url.URL) {
	var getParamsURI string
	var _val string
	for key, val := range httpRequest.GetParams {
		_val = url.QueryEscape(val)
		if getParamsURI == "" {
			getParamsURI = fmt.Sprintf("%s=%s", key, _val)
		} else {
			getParamsURI = fmt.Sprintf("%s&%s=%s", getParamsURI, key, _val)
		}
	}
	requestUrl := fmt.Sprintf("%s?%s", httpRequest.Url, getParamsURI)
	uri, err := url.Parse(requestUrl)

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

func (httpRequest *HTTPRequest) Patch() (body string, err error) {
	httpRequest.Method = "PATCH"
	body, err = httpRequest.httpResponseBody()
	return
}

func (httpRequest *HTTPRequest) Delete() (body string, err error) {
	httpRequest.Method = "DELETE"
	body, err = httpRequest.httpResponseBody()
	return
}

func (httpRequest *HTTPRequest) Head() (*http.Response, error) {
	httpRequest.Method = "HEAD"
	return httpRequest.httpResponse()
}

func (httpRequest *HTTPRequest) Options() (*http.Response, error) {
	httpRequest.Method = "OPTIONS"
	return httpRequest.httpResponse()
}
