package golgoquery

import (
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

func Goquery(url string, goquerySelector string) (domNode *goquery.Selection) {
	doc := GoqueryDocument(url)

	domNode = doc.Find(goquerySelector)
	return
}

func (results *GoqueryResults) GoqueryAttrsFrom(domNodes *goquery.Selection, attr string) {
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
			s_html, _ := s.Html()
			log.Printf("[warn] %s\n", s_html)
		}

		results.Results[i] = attrValue
	})
	return
}

func GoqueryHrefsFrom(url string, goquerySelector string) (results GoqueryResults) {
	results.GoqueryAttrsFrom(Goquery(url, goquerySelector), "href")
	return
}

func GoqueryTextFrom(url string, goquerySelector string) (results GoqueryResults) {
	results.GoqueryAttrsFrom(Goquery(url, goquerySelector), "text")
	return
}

func GoqueryAttrsFromParents(url string, selectors []string, attr string) (results GoqueryResults) {
	var domNodes *goquery.Selection
	last_idx := len(selectors) - 1
	for idx, selector := range selectors {
		if selector == ".." {
			domNodes = domNodes.Parent()
		} else {
			domNodes = Goquery(url, selector)
		}
		if idx == last_idx {
			results.GoqueryAttrsFrom(domNodes, attr)
		}
	}
	return
}

func GoqueryHrefsFromParents(url string, selectors []string) (results GoqueryResults) {
	results = GoqueryAttrsFromParents(url, selectors, "href")
	return
}

func GoqueryTextFromParents(url string, selectors []string) (results GoqueryResults) {
	results = GoqueryAttrsFromParents(url, selectors, "text")
	return
}
