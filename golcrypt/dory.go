package golcrypt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/abhishekkr/gol/golerror"
	"github.com/abhishekkr/gol/golhttpclient"
	"github.com/abhishekkr/gol/golrandom"
)

type Dory struct {
	BaseUrl       string
	Backend       string //local-auth
	Key           string
	Value         []byte
	Token         string
	KeyTTL        int //seconds, usable in cache mode only i.e. when persist is false
	Persist       bool
	ReadNotDelete bool
}

func (dory *Dory) userBackend() string {
	backend := dory.Backend
	if backend == "" {
		backend = "local-auth"
	}
	return backend
}

func (dory *Dory) adminBackend() string {
	if dory.Persist {
		return "disk"
	}
	return "cache"
}

func (dory *Dory) httpUserUrl(request *golhttpclient.HTTPRequest) {
	request.Url = fmt.Sprintf("%s/%s/%s", dory.BaseUrl, dory.userBackend(), dory.Key)
}

func (dory *Dory) httpAdminUrl(request *golhttpclient.HTTPRequest) {
	request.Url = fmt.Sprintf("%s/admin/store/%s", dory.BaseUrl, dory.adminBackend())
}

func (dory *Dory) httpUserHeaders(request *golhttpclient.HTTPRequest) {
	request.HTTPHeaders = map[string]string{
		"X-DORY-TOKEN": dory.Token,
	}
}

func (dory *Dory) httpAdminHeaders(request *golhttpclient.HTTPRequest) {
	request.HTTPHeaders = map[string]string{
		"X-DORY-ADMIN-TOKEN": dory.Token,
	}
}

func (dory *Dory) httpParams(request *golhttpclient.HTTPRequest) {
	ttlsecond := strconv.Itoa(dory.KeyTTL)
	if ttlsecond == "" {
		ttlsecond = "300"
	}
	request.GetParams = map[string]string{
		"keep":      fmt.Sprintf("%t", dory.ReadNotDelete),
		"persist":   fmt.Sprintf("%t", dory.Persist),
		"ttlsecond": ttlsecond,
	}
}

func (dory *Dory) httpUserRequest() golhttpclient.HTTPRequest {
	request := golhttpclient.HTTPRequest{}
	dory.httpUserUrl(&request)
	dory.httpUserHeaders(&request)
	dory.httpParams(&request)
	return request
}

func (dory *Dory) httpAdminRequest() golhttpclient.HTTPRequest {
	request := golhttpclient.HTTPRequest{}
	dory.httpAdminUrl(&request)
	dory.httpAdminHeaders(&request)
	dory.httpParams(&request)
	return request
}

func (dory *Dory) Set() (err error) {
	if dory.Key == "" {
		dory.Key = fmt.Sprintf("dory-%s", golrandom.Token(10))
	}

	request := golhttpclient.HTTPRequest{}
	dory.httpUserUrl(&request)
	request.Body = bytes.NewBuffer(dory.Value)

	dory.Token, err = request.Post()
	return
}

func (dory *Dory) ShareSecret(value []byte) (err error) {
	dory.Value = value
	err = dory.Set()
	return
}

func (dory *Dory) ShareSecretFromFile(filepath string) (err error) {
	requestBody, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	err = dory.ShareSecret(requestBody)
	return
}

func (dory *Dory) Get() (err error) {
	if dory.BaseUrl == "" {
		err = golerror.Error(123, "dory url can't be empty")
		return
	}
	if dory.Key == "" || dory.Token == "" {
		err = golerror.Error(123, "key or token can't be empty")
		return
	}

	request := dory.httpUserRequest()

	response, err := request.Get()
	dory.Value = []byte(response)
	return
}

func (dory *Dory) FetchSecret() (value []byte, err error) {
	err = dory.Get()
	value = dory.Value
	return
}

func (dory *Dory) RefreshSecret() (value []byte, err error) {
	readNotDelete := dory.ReadNotDelete
	dory.ReadNotDelete = true
	value, err = dory.FetchSecret()
	dory.ReadNotDelete = readNotDelete
	return
}

func (dory *Dory) Del() (err error) {
	if dory.BaseUrl == "" {
		err = golerror.Error(123, "dory url can't be empty")
		return
	}
	if dory.Key == "" || dory.Token == "" {
		err = golerror.Error(123, "key and token required to purge")
		return
	}

	request := dory.httpUserRequest()

	response, err := request.Get()
	dory.Value = []byte(response)
	return
}

func (dory *Dory) PurgeSecret() (err error) {
	return dory.Del()
}

func (dory *Dory) PurgeAll() (err error) {
	if dory.BaseUrl == "" {
		err = golerror.Error(123, "dory url can't be empty")
		return
	}
	if dory.Token == "" {
		err = golerror.Error(123, "admin token required to purge")
		return
	}

	request := dory.httpAdminRequest()

	response, err := request.Delete()
	dory.Value = []byte(response)
	return
}

func (dory *Dory) List() (err error) {
	if dory.BaseUrl == "" {
		err = golerror.Error(123, "dory url can't be empty")
		return
	}
	if dory.Token == "" {
		err = golerror.Error(123, "admin token required to purge")
		return
	}

	request := dory.httpAdminRequest()

	response, err := request.Get()
	dory.Value = []byte(response)
	return
}

func (dory *Dory) Ping() (err error) {
	if dory.BaseUrl == "" {
		err = golerror.Error(123, "dory url can't be empty")
		return
	}
	request := golhttpclient.HTTPRequest{}
	request.Url = fmt.Sprintf("%s/ping", dory.BaseUrl)

	response, err := request.Get()
	dory.Value = []byte(response)
	return
}
