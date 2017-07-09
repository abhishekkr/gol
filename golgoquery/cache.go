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
	ReloadCache bool
	CacheDir    = "/tmp/.tune.cli"
	UserAgent   = "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:53.0) Gecko/20100101 Firefox/53.0"
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

func cleanCache(cachePath string) {
	var err = os.Remove(cachePath)
	if err != nil {
		log.Printf("[warn] failed cleaning %s", cachePath)
	}
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

	if ReloadCache {
		log.Println("[warn] cleaning cache ", cachePath)
		cleanCache(cachePath)
	}

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		headers := map[string]string{
			"User-Agent": UserAgent,
		}
		log.Println("[info] fetching ", url)
		body, err := golhttpclient.HttpGet(url, map[string]string{}, headers)
		if err != nil {
			log.Fatalf("not able to fetch %s", url)
		}
		createCache(cachePath, body)
	}

	cache := readCache(cachePath)

	return goquery.NewDocumentFromReader(strings.NewReader(cache))
}
