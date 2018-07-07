package golgoquery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

/*
Action type is for the actual functions that will process query's results and return same or processed result.
*/
type Action func([]GoqueryResults, int) ([]GoqueryResults, error)

type QueryDSL struct {
	URIFlows [](*URIFlow) `json:"uriflows"`
}

type URIFlow struct {
	URI        URI            `json:"uri"`
	QueryFlows [](*QueryFlow) `json:"queryflows"`
}

type URI string

type QueryFlow struct {
	ResultsList []GoqueryResults `json:"content"`
	Queries     [](*Query)       `json:"queries"`
}

type DOMSelection *goquery.Document

type Query struct {
	Selector   []string `json:"selector"`
	Attribute  string   `json:"attribute"`
	ActionName string   `json:"action"`
}

/*
QueryDSL.Proc processes QueryDSL DSL block.
*/
func (queryDSL *QueryDSL) Proc(actionmap map[string]Action) {
	for _, uriflow := range queryDSL.URIFlows {
		uriflow.Proc(actionmap)
	}
}

/*
URIFlow.Proc processes URIFlow DSL block.
*/
func (uriFlow *URIFlow) Proc(actionmap map[string]Action) {
	for _, queryflow := range uriFlow.QueryFlows {
		queryflow.Proc(actionmap, uriFlow.URI)
	}
}

/*
QueryFlow.Proc processes QueryFlow DSL block.
*/
func (queryFlow *QueryFlow) Proc(actionmap map[string]Action, uri URI) {
	queryFlow.ResultsList = make([]GoqueryResults, len(queryFlow.Queries))
	for idx, q := range queryFlow.Queries {
		var err error
		queryFlow.ResultsList[idx], err = GoqueryAttrsFromParents(string(uri),
			q.Selector,
			q.Attribute)
		if err != nil {
			fmt.Println(err)
			return
		}

		queryFlow.ResultsList, err = actionmap[q.ActionName](queryFlow.ResultsList, idx)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/*
FromJson can be passed QueryDSL representing JSON blob and map of action name to function, to process the DSL.
*/
func FromJson(jsonBlob []byte, actionmap map[string]Action) (qDSL QueryDSL, err error) {
	err = json.Unmarshal(jsonBlob, &qDSL)
	if err != nil {
		return
	}
	qDSL.Proc(actionmap)
	return
}

/*
FromFile can be passed file with JSON blob representing QueryDSL and map of action name to function, to process the DSL.
*/
func FromFile(path string, actionmap map[string]Action) (QueryDSL, error) {
	fileblob, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return FromJson(fileblob, actionmap)
}

/*
Debug is a sample built-in action. That can be passed as action for query/queries to print results during development phase.
*/
func Debug(resultsList []GoqueryResults, idx int) ([]GoqueryResults, error) {
	for _, result := range resultsList[idx].Results {
		fmt.Println(result)
	}
	return resultsList, nil
}
