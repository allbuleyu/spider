package main

import (
	"flag"
	"fmt"
	"github.com/allbuleyu/spider/spider"
	"github.com/astaxie/beego/logs"
	"net/http"

	"runtime"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
	Refer     = ""
)

var (
	startPage = flag.String("s", "http://www.douban.com/group/haixiuzu/discussion", "douban group start page")
)

type container int

type containers struct {
	container
	Len int
}

func (c *containers) Iter() <-chan container {
	item := make(chan container)

	go func() {
		for i := 0; i < c.Len; i++ {
			item <- c.container
			// item = <-c.container  invalid operation: <-c.container (receive from non-chan type container)
		}
	}()

	return item
}

func watchChanRes() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 123
	}()

	return ch
}

func main() {
	fmt.Println(runtime.NumCPU(), runtime.NumCgoCall())
	c := new(containers)

	var con container = 1

	c.Len = 10

	for i := 0; i < 10; i++ {
		c.container = con
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-c.Iter())
	}
	// fmt.Println(<-watchChanRes())
}

func t() {
	flag.Parse()
	url := *startPage

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.Error(err)
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Referer", Refer)

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		logs.Error(err)
	}

	s, err := spider.CreateSpiderFromResponse(resp)
	if err != nil {
		logs.Error(err)
	}

	rs, _ := s.GetAttr("div.grid-16-8.clearfix>div.article>div>table.olt>tbody>tr>td.title>a", "href")

	fmt.Println(rs)
}
