package main

import (
	"fmt"

	golconfig "github.com/abhishekkr/gol/golconfig"
)

type dictOfDict map[string]map[string]string

func main() {
	var conf dictOfDict
	conf = make(dictOfDict)

	json := golconfig.GetConfig("json")

	dat := "{\"AAA\": {\"abc\": \"123\", \"xyz\": \"456\"}}"

	json.Config(dat, &conf)
	fmt.Println(conf)
}
