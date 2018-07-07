package main

import (
	"fmt"
	"log"

	golgoquery "github.com/abhishekkr/gol/golgoquery"
)

func Dummy(resultsList []golgoquery.GoqueryResults, idx int) ([]golgoquery.GoqueryResults, error) {
	for _, result := range resultsList[idx].Results {
		fmt.Println(result)
	}
	return resultsList, nil
}

var (
	selector  = []string{"h1"}
	attribute = "text"
	actionmap = map[string]golgoquery.Action{"Dummy": Dummy}

	queries = [](*golgoquery.Query){
		&golgoquery.Query{
			Selector:   selector,
			Attribute:  attribute,
			ActionName: "Dummy",
		},
	}

	queryFlow = golgoquery.QueryFlow{Queries: queries}

	uriFlow = golgoquery.URIFlow{
		URI:        golgoquery.URI("http://example.com"),
		QueryFlows: [](*golgoquery.QueryFlow){&queryFlow},
	}

	workflow = golgoquery.QueryDSL{
		URIFlows: [](*golgoquery.URIFlow){&uriFlow},
	}

	jsonblobx = `
{ "uriflows": [
    { "uri": "http://example.com/",
      "queryflows": [
        { "queries": [
				{ "selector": ["a"], "attribute": "text", "action": "Dummy" }
			] }
      ] }
  ] }`

	jsonbloby = `
{ "uriflows": [
    { "uri": "http://example.com/",
      "queryflows": [ {
			"queries": [
				{ "selector": ["div", "p"], "attribute": "text", "action": "Dummy" },
				{ "selector": ["div", "a"], "attribute": "href", "action": "Dummy" }
			] }
      ] }
  ] }`
)

func main() {
	workflow.Proc(actionmap)
	log.Printf("%v", workflow.URIFlows[0].QueryFlows[0].ResultsList)
	qDSL, err := golgoquery.FromJson([]byte(jsonblobx), actionmap)
	if err == nil {
		log.Printf("%v", qDSL.URIFlows[0].QueryFlows[0].ResultsList)
	} else {
		log.Println(err)
	}
	qDSL, err = golgoquery.FromJson([]byte(jsonbloby), actionmap)
	if err == nil {
		log.Printf("%v", qDSL.URIFlows[0].QueryFlows[0].ResultsList)
	} else {
		log.Println(err)
	}
	fmt.Println("~finito")
}
