package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"sync"
)

// Spider
type Spider struct {
	Url string // page that spider would deal with
	doc *goquery.Document
}

// Start spider
func CreateSpiderFromUrl(url string) (*Spider, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, fmt.Errorf("url %s, error %s", url, err)
	}
	return &Spider{Url: url, doc: doc}, nil
}

func CreateSpiderFromResponse(r *http.Response) (*Spider, error) {
	doc, err := goquery.NewDocumentFromResponse(r)
	if err != nil {
		return nil, fmt.Errorf("error %s", err)
	}
	return &Spider{doc: doc}, nil
}
