package golcrypt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/abhishekkr/gol/golhttpclient"
	"github.com/abhishekkr/gol/golrandom"
)

type Dory struct {
	BaseUrl string
	Backend string //local-auth
	Key     string
	Token   string
}

func (dory *Dory) ShareSecret(value []byte) (err error) {
	request := golhttpclient.HTTPRequest{}

	backend := dory.Backend
	if backend == "" {
		backend = "local-auth"
	}

	key := dory.Key
	if key == "" {
		key = fmt.Sprintf("dory-%s", golrandom.Token(10))
	}

	request.Url = fmt.Sprintf("%s/%s/%s", dory.BaseUrl, backend, key)
	request.Body = bytes.NewBuffer(value)

	dory.Token, err = request.Post()
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

func (dory *Dory) FetchSecret() (value []byte, err error) {
	request := golhttpclient.HTTPRequest{}

	backend := dory.Backend
	if backend == "" {
		backend = "local-auth"
	}

	if dory.Key == "" || dory.Token == "" {
		log.Fatalln("key and token need to be provided to fetch")
	}

	request.Url = fmt.Sprintf("%s/%s/%s", dory.BaseUrl, backend, dory.Key)

	request.HTTPHeaders = map[string]string{
		"X-DORY-TOKEN": dory.Token,
	}

	dory.Token, err = request.Get()
	return
}

func (dory *Dory) RefreshSecret() (value []byte, err error) {
	request := golhttpclient.HTTPRequest{}

	backend := dory.Backend
	if backend == "" {
		backend = "local-auth"
	}

	if dory.Key == "" || dory.Token == "" {
		log.Fatalln("key and token need to be provided to fetch")
	}

	request.Url = fmt.Sprintf("%s/%s/%s", dory.BaseUrl, backend, dory.Key)

	request.GetParams = map[string]string{
		"keep": "true",
	}

	request.HTTPHeaders = map[string]string{
		"X-DORY-TOKEN": dory.Token,
	}

	dory.Token, err = request.Get()
	return
}

func (dory *Dory) PurgeSecret() (err error) {
	request := golhttpclient.HTTPRequest{}

	backend := dory.Backend
	if backend == "" {
		backend = "local-auth"
	}

	key := dory.Key
	if key == "" {
		log.Fatalln("no key provided to purge")
	}

	request.Url = fmt.Sprintf("%s/%s/%s", dory.BaseUrl, backend, dory.Key)

	request.HTTPHeaders = map[string]string{
		"X-DORY-TOKEN": dory.Token,
	}

	dory.Token, err = request.Delete()
	return
}
