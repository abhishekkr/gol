package golgoquery

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/abhishekkr/gol/golhttpclient"
)

var (
	CacheDir string
)

func createCache(cachePath string, cacheContent string) {
	cacheDir := path.Dir(cachePath)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(CacheDir, 0777)
	}
}

func readCache(cachePath string) string {
	dat, err := ioutil.ReadFile(cachePath)
	return string(dat)
}

func urlToFilename(url string) string {
	var replacer = strings.NewReplacer(" ", "-",
		"\t", "-",
		"%", "-",
		"?", "-",
		"=", "-",
		"\\", "-",
		"/", "-",
	)
	return replacer.Replace(url)
}

func CacheUrl(url string) (*goquery.Document, error) {
	urlFile = urlToFilename(url)
	if CacheDir == "" {
		if _, err := os.Stat("/tmp"); os.IsNotExist(err) {
			log.Fatalf("cannot find default CacheDir(%s), default cachedir is not created", CacheDir)
		}
		CacheDir = "/tmp"
	}
	cachePath := fmt.Sprintf("%s%s%s", CacheDir, filepath.Separator, urlFile)

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		body, err := golhttpclient.HttpGet(url)
		if err != nil {
			log.Fatalf("not able to fetch %s", url)
		}
		createCache(cachePath, body)
	}

	cache := readCache(cachePath)

	utfBody, err := iconv.NewReader(cache, charset, "utf-8")
	if err != nil {
		log.Fatalf("not able to read data of %s", url)
	}

	return goquery.NewDocumentFromReader(strings.NewReader(utfBody))
}
