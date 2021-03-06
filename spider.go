package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://licaizone.com/")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".slideshow #carousel-example-generic .carousel-inner .item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, err := s.Find("img").Attr("alt")
		// title := s.Find("i").Text()
		fmt.Printf("Review %d: %s\n", i, band, err)
	})
}

type OnePiecePic struct {
	Id          string `akke`
	Group_id    string
	Name        string
	Path        string
	Create_time string
	Is_delete   string
}

type data map[string]interface{}

func main() {
	// ExampleScrape()
	db, err := sql.Open("mysql", "admin:Dream1tPossible@tcp(114.215.154.110:3306)/first-go?charset=utf8")

	defer db.Close()

	rows, err := db.Query("select * from one_piece_pic where id > ?", 1)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	scanArgs := make([]interface{}, len(columns))

	op := new(OnePiecePic)
	ops := make([]OnePiecePic, 0)

	t := reflect.TypeOf(op).Elem()
	v := reflect.ValueOf(op).Elem()

	if len(columns) < t.NumField() {
		panic("columns length is less than struct length")
	}

	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		fieldName := tf.Name
		// fmt.Printf("%#v %v \n", vf.Addr(), vf.Addr())
		for ii := range columns {
			if strings.ToLower(columns[ii]) == strings.ToLower(fieldName) && vf.CanAddr() {
				// fmt.Println(vf.Addr())
				scanArgs[ii] = vf.Addr()

				// fmt.Printf("%#v %v %d \n", scanArgs[ii], scanArgs[ii], ii)
			}
		}
	}
	// fmt.Println(scanArgs, 111, ops)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		ops = append(ops, *op)
	}

	fmt.Println(ops)
}

func test() {
	db, err := sql.Open("mysql", "admin:Dream1tPossible@tcp(114.215.154.110:3306)/first-go?charset=utf8")

	defer db.Close()

	rows, err := db.Query("select * from one_piece_pic where id > ?", 1)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	// values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(columns))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	op := new(OnePiecePic)
	ops := make([]OnePiecePic, 0)
	// 这里可以优化为一个循环,兼容各个不同的数据类型
	scanArgs[0] = &op.Id
	scanArgs[1] = &op.Group_id
	scanArgs[2] = &op.Name
	scanArgs[3] = &op.Path
	scanArgs[4] = &op.Create_time
	scanArgs[5] = &op.Is_delete

	fmt.Println(scanArgs)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		ops = append(ops, *op)
	}

	fmt.Println(scanArgs, ops)
}
