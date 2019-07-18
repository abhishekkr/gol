package golhttpclient

import (
	"bufio"
	"bytes"
	"strings"
)

const (
	initSpec    = 0
	methodSpec  = 1
	headersSpec = 2
	bodySpec    = 3
)

type HTTPSpec struct {
	Client *HTTPRequest

	Spec string
}

func (spec *HTTPSpec) Parse() {
	state := 0
	var body bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(spec.Spec))
	spec.Client = &HTTPRequest{
		GetParams:   map[string]string{},
		HTTPHeaders: map[string]string{},
	}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if state == initSpec {
				state = methodSpec
			} else if state == headersSpec {
				state = bodySpec
			}
			continue
		} else if state <= methodSpec {
			spec.parseMethodURLAndParams(txt)
			state = headersSpec
		} else if state == headersSpec {
			spec.parseHeader(txt)
		} else if state == bodySpec {
			body.WriteString(txt)
		}
	}
	spec.Client.Body = &body
}

func (spec *HTTPSpec) parseMethodURLAndParams(txt string) {
	txtParts := strings.Split(txt, " ")
	lastIndex := len(txtParts) - 1
	spec.Client.Method = txtParts[0]
	spec.Client.Protocol = txtParts[lastIndex]

	_url := strings.Join(txtParts[1:lastIndex], "")
	_url = strings.Replace(_url, "&amp; ", "&", -1)
	_urlParts := strings.Split(_url, "?")
	spec.Client.Url = _urlParts[0]
	if len(_urlParts) == 1 {
		return
	}
	for _, param := range strings.Split(_urlParts[1], "&") {
		getKeyVal := strings.Split(param, "=")
		if len(getKeyVal) == 1 {
			spec.Client.GetParams[getKeyVal[0]] = ""
			continue
		}
		spec.Client.GetParams[getKeyVal[0]] = strings.Join(getKeyVal[1:(len(getKeyVal)-1)], "=")
	}
}

func (spec *HTTPSpec) parseHeader(txt string) {
	txtParts := strings.Split(txt, ":")
	headerName := strings.TrimSpace(txtParts[0])
	headerValue := strings.TrimSpace(
		strings.Join(txtParts[1:len(txtParts)], ":"),
	)
	spec.Client.HTTPHeaders[headerName] = headerValue
}
