package main

import (
	// "flag"
	"fmt"

	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"sync"

	"time"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
	Refer     = ""
)

var (
	// 这个地址是个ajax页面,返回json数组
	// startPage = flag.String("s", "https://movie.douban.com/j/search_subjects?type=movie&tag=豆瓣高分&sort=recommend&page_limit=20&page_start=0", "douban group start page")

	startPage = "https://movie.douban.com/j/search_subjects?type=movie&tag=日本&sort=recommend&page_limit=20&page_start=%d"
)

type Sub struct {
	Subjects []Movie
}

type Movie struct {
	Id     string
	Rate   string
	Title  string
	Url    string
	Cover  string
	Is_new bool
}

func main() {
	startTime := time.Now()

	i := 0
	for {
		url := fmt.Sprintf(startPage, i*20)
		isDone := MovieMain(url)

		if isDone {
			break
		}

		i++

		time.Sleep(time.Second * 3)
	}

	elapsed := time.Since(startTime)

	fmt.Println(elapsed)
}

func MovieMain(url string) bool {
	startTime := time.Now()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.Error(err)
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Referer", Refer)

	client := new(http.Client)

	logs.SetLogger("file", `{"filename":"test.log"}`)

	// resp, err := client.Do(req)
	resp, err := client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		logs.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	movies := new(Sub)

	json.Unmarshal(body, movies)

	if len(movies.Subjects) == 0 {
		return true
	}

	db, err := sql.Open("mysql", "admin:Dream1tPossible@tcp(114.215.154.110:3306)/first-go?charset=utf8")

	defer db.Close()

	var wg sync.WaitGroup

	var q string = ""
	for _, m := range movies.Subjects {
		wg.Add(1)
		// db_id, _ := strconv.ParseInt(movie.Id, 10, 0)
		// rate, _ := strconv.ParseFloat(movie.Rate, 0)
		// q = fmt.Sprintf("insert into movie (rate, name, db_id, is_new, cover, url) VALUES (%f, '%s', %d, %d, '%s', '%s')", rate, movie.Title, db_id, is_new, movie.Cover, movie.Url)

		go func(movie Movie) {
			defer wg.Done()

			is_new := BooleToInt(movie.Is_new)

			q = fmt.Sprintf("insert into movie (rate, name, db_id, is_new, cover, url, group_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
			res, err := db.Exec(q, movie.Rate, movie.Title, movie.Id, is_new, movie.Cover, movie.Url, 5)

			if err != nil {
				logs.Error("this err is %s", err)
			}

			if err == nil {
				fmt.Println(res.LastInsertId())
				fmt.Println("---------------------")
			}
		}(m)

	}

	wg.Wait()
	elapsed := time.Since(startTime)

	logs.Info("run time is %s \n\n\n\n", elapsed)

	return false
}

func BooleToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
