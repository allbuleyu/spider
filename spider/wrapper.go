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

func (s *Spider) GetAttr(rule, attr string) ([]string, error) {
	var (
		res = make([]string, 0) //for leaf
		wg  sync.WaitGroup
		mu  sync.Mutex
	)

	s.doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			attr, ok := sl.Attr(attr)
			if ok {
				mu.Lock()
				res = append(res, attr)
				mu.Unlock()
			}
		}()
	})
	wg.Wait()
	return res, nil
}

// 获取爬到的HTML的属性
func (s *Spider) GetAttrs(rule, attr string) ([]string, error) {
	var (
		res = make([]string, 0)
	)

	s.doc.Find(rule).Each(func(i int, sl *goquery.Selection) {
		attr, ok := sl.Attr(attr)

		if ok {
			res = append(res, attr)
		}
	})

	return res, nil
}

func (s *Spider) GetJsonData() {

}
