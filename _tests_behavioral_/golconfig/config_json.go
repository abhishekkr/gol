package main

import (
	"fmt"

	golassert "github.com/abhishekkr/gol/golassert"
	golconfig "github.com/abhishekkr/gol/golconfig"
)

type dictOfDict map[string]map[string]string

func main() {
	var conf dictOfDict
	conf = make(dictOfDict)

	json := golconfig.GetConfigurator("json")

	dat := "{\"AAA\": {\"abc\": \"123\", \"xyz\": \"456\"}}"

	json.Unmarshal(dat, &conf)
	if conf["AAA"]["abc"] != "123" && conf["AAA"]["xyz"] != "456" {
		fmt.Println("FAILED!")
	}
	golassert.AssertEqual(conf["AAA"]["abc"], "123")
	golassert.AssertEqual(conf["AAA"]["xyz"], "456")
	fmt.Println(conf)
}
