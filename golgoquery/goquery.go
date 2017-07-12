package golgoquery

import (
	"errors"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var (
	CacheGoquery bool
)

type GoqueryResults struct {
	Results []string
}

func GoqueryDocument(url string) *goquery.Document {
	var doc *goquery.Document
	var err error

	if CacheGoquery == true {
		doc, err = CacheUrl(url)
	} else {
		doc, err = goquery.NewDocument(url)
	}
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GoqueryFromDocument(doc *goquery.Document, goquerySelector string) (domNode *goquery.Selection) {
	domNode = doc.Find(goquerySelector)
	return
}

func Goquery(url string, goquerySelector string) (domNode *goquery.Selection) {
	doc := GoqueryDocument(url)

	domNode = GoqueryFromDocument(doc, goquerySelector)
	return
}

func (results *GoqueryResults) GoqueryAttrsFrom(domNodes *goquery.Selection, attr string) (err error) {
	results.Results = make([]string, domNodes.Size())
	domNodes.Each(func(i int, s *goquery.Selection) {
		var attrValue string
		var attrPresent bool
		if attr == "text" {
			attrValue = s.Text()
			attrPresent = (attrValue != "")
		} else {
			attrValue, attrPresent = s.Attr(attr)
		}
		if !attrPresent {
			if err == nil {
				err = errors.New(fmt.Sprintf("%s", attr))
			} else {
				err = errors.New(fmt.Sprintf("%s | %s", err.Error(), attr))
			}
		}

		results.Results[i] = attrValue
	})
	return nil
}

func GoqueryHrefsFrom(url string, goquerySelector string) (results GoqueryResults, err error) {
	err = results.GoqueryAttrsFrom(Goquery(url, goquerySelector), "href")
	return
}

func GoqueryTextFrom(url string, goquerySelector string) (results GoqueryResults, err error) {
	err = results.GoqueryAttrsFrom(Goquery(url, goquerySelector), "text")
	return
}

func GoqueryAttrsFromParents(url string, selectors []string, attr string) (results GoqueryResults, err error) {
	domNodes := Goquery(url, selectors[0])
	for idx := 1; idx < len(selectors); idx++ {
		if selectors[idx] == ".." {
			domNodes = domNodes.Parent()
		} else {
			domNodes.Find(selectors[idx])
		}
	}
	err = results.GoqueryAttrsFrom(domNodes, attr)
	return
}

func GoqueryHrefsFromParents(url string, selectors []string) (results GoqueryResults, err error) {
	results, err = GoqueryAttrsFromParents(url, selectors, "href")
	return
}

func GoqueryTextFromParents(url string, selectors []string) (results GoqueryResults, err error) {
	results, err = GoqueryAttrsFromParents(url, selectors, "text")
	return
}
