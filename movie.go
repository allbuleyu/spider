package main

import (
	"flag"
	"fmt"
	// "github.com/allbuleyu/spider/spider"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
	Refer     = ""
)

var (
	// 这个地址是个ajax页面,返回json数组
	startPage = flag.String("s", "https://movie.douban.com/j/search_subjects?type=movie&tag=豆瓣高分&sort=recommend&page_limit=20&page_start=0", "douban group start page")
)

type Sub struct {
	Subjects []Movie
}

type Movie struct {
	Rate   string
	Title  string
	Url    string
	Cover  string
	Is_new bool
}

func main() {
	flag.Parse()
	url := *startPage
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.Error(err)
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Referer", Refer)

	client := new(http.Client)

	// resp, err := client.Do(req)
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		logs.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	movie := new(Sub)

	json.Unmarshal(body, movie)

	fmt.Println(movie)
}
