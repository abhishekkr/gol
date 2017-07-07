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
	CacheDir = "/tmp/.tune.cli"
)

func createCache(cachePath string, cacheContent string) {
	cacheDir := path.Dir(cachePath)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		log.Printf("[warn] cannot find default CacheDir(%s), creating cachedir", CacheDir)
		os.Mkdir(CacheDir, 0777)
	}

	err := ioutil.WriteFile(cachePath, []byte(cacheContent), 0644)
	if err != nil {
		log.Printf("[warn] cannot create cache (%s), creating cachedir", cachePath)
	}
}

func readCache(cachePath string) string {
	dat, err := ioutil.ReadFile(cachePath)
	if err != nil {
		log.Printf("[warn] missing cache file %s", cachePath)
	}
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
		":", "-",
	)
	return replacer.Replace(url)
}

func CacheUrl(url string) (*goquery.Document, error) {
	urlFile := urlToFilename(url)
	cachePath := fmt.Sprintf("%s%s%s", CacheDir, string(filepath.Separator), urlFile)

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		body, err := golhttpclient.HttpGet(url, map[string]string{}, map[string]string{})
		if err != nil {
			log.Fatalf("not able to fetch %s", url)
		}
		createCache(cachePath, body)
	}

	cache := readCache(cachePath)

	return goquery.NewDocumentFromReader(strings.NewReader(cache))
}
