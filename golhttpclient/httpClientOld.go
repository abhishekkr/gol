package golhttpclient

import "net/http"

func httpResponseBody(httpMethod string, baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	httpRequest := HTTPRequest{
		Method:      httpMethod,
		Url:         baseURL,
		GetParams:   getParams,
		HTTPHeaders: httpHeaders,
	}

	return httpRequest.httpResponseBody()
}

func Http(httpMethod string, baseURL string, getParams map[string]string, httpHeaders map[string]string) (resp *http.Response, err error) {
	httpRequest := HTTPRequest{
		Method:      httpMethod,
		Url:         baseURL,
		GetParams:   getParams,
		HTTPHeaders: httpHeaders,
	}

	return httpRequest.httpResponse()
}

func HttpGet(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	body, err = httpResponseBody("GET", baseURL, getParams, httpHeaders)
	return
}

func HttpPut(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	// need to handle PUT body content
	body, err = httpResponseBody("PUT", baseURL, getParams, httpHeaders)
	return
}

func HttpPost(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	// need to handle POST body content
	body, err = httpResponseBody("POST", baseURL, getParams, httpHeaders)
	return
}

func HttpDelete(baseURL string, getParams map[string]string, httpHeaders map[string]string) (body string, err error) {
	body, err = httpResponseBody("DELETE", baseURL, getParams, httpHeaders)
	return
}
