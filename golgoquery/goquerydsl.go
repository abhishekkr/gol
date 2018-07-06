package golgoquery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

/*
{ "uriflows": [
  { "uri": "http://example.com/",
    "queryflows": [
	  {
	    "queries": [
	      {
			"selector": ["div", "p"],
			"attribute": "text",
			"action": "Dummy"
		  },
	      {
			"selector": ["div", "a"],
			"attribute": "href",
			"action": "Dummy"
		  }
	    ]
      }
	]
  }
]}

above DSL
* fetchs "uri"
* picks every query from "queryflows"
> * first "query", pulls innerHTML text from "div.p" selector element; passes list of results to Method mapped to "Dummy"
> * then next "query", pulls href from "div.a" selector element; passes list of results to Method mapped to "Dummy"

*/

type Action func(GoqueryResults) (GoqueryResults, error)

type QueryDSL struct {
	URIFlows []URIFlow `json:"uriflows"`
}

type URIFlow struct {
	URI        URI         `json:"uri"`
	QueryFlows []QueryFlow `json:"queryflows"`
}

type URI string

type QueryFlow struct {
	Content GoqueryResults `json:"content"`
	Queries []Query        `json:"queries"`
}

type DOMSelection *goquery.Document

type Query struct {
	Selector   []string `json:"selector"`
	Attribute  string   `json:"attribute"`
	ActionName string   `json:"action"`
}

func (queryDSL QueryDSL) Proc(actionmap map[string]Action) {
	for _, uri := range queryDSL.URIFlows {
		uri.Proc(actionmap)
	}
}

func (uriFlow URIFlow) Proc(actionmap map[string]Action) {
	for _, query := range uriFlow.QueryFlows {
		query.Proc(actionmap, uriFlow.URI)
	}
}

func (queryFlow QueryFlow) Proc(actionmap map[string]Action, uri URI) {
	for _, q := range queryFlow.Queries {
		var err error
		queryFlow.Content, err = GoqueryAttrsFromParents(string(uri),
			q.Selector,
			q.Attribute)
		if err != nil {
			fmt.Println(err)
			return
		}

		queryFlow.Content, err = actionmap[q.ActionName](queryFlow.Content)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/*
func SelectionText(doc *goquery.Document, matcher string) (txt string) {
	doc.Find(matcher).Each(func(i int, htmlSel *goquery.Selection) {
		txt += htmlSel.Text()
	})
	return
}
*/

func FromJson(jsonBlob []byte, actionmap map[string]Action) error {
	var qDSL QueryDSL
	err := json.Unmarshal(jsonBlob, &qDSL)
	if err != nil {
		return err
	}
	qDSL.Proc(actionmap)
	return nil
}

func FromFile(path string, actionmap map[string]Action) error {
	fileblob, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return FromJson(fileblob, actionmap)
}
