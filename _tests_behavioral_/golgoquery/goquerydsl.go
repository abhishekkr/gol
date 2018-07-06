package main

import (
	"fmt"

	golgoquery "github.com/abhishekkr/gol/golgoquery"
)

func Dummy(results golgoquery.GoqueryResults) (golgoquery.GoqueryResults, error) {
	for _, result := range results.Results {
		fmt.Println(result)
	}
	return results, nil
}

var (
	selector  = []string{"h1"}
	attribute = "text"
	actionmap = map[string]golgoquery.Action{"Dummy": Dummy}

	queries = []golgoquery.Query{
		golgoquery.Query{
			Selector:   selector,
			Attribute:  attribute,
			ActionName: "Dummy",
		},
	}

	queryFlow = golgoquery.QueryFlow{Queries: queries}

	uriFlow = golgoquery.URIFlow{
		URI:        golgoquery.URI("http://example.com"),
		QueryFlows: []golgoquery.QueryFlow{queryFlow},
	}

	workflow = golgoquery.QueryDSL{
		URIFlows: []golgoquery.URIFlow{uriFlow},
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
	golgoquery.FromJson([]byte(jsonblobx), actionmap)
	golgoquery.FromJson([]byte(jsonbloby), actionmap)
	fmt.Println("~finito")
}
