package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	Id          string
	Group_id    string
	Name        string
	Path        string
	Create_time string
	Is_delete   string
}

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

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		ops = append(ops, *op)
	}

	fmt.Println(scanArgs, ops)
}

func test() {
	// var op OnePiecePic
	// ops := make([]OnePiecePic, 0)
	// // panic(1)

	// // Fetch rows
	// for rows.Next() {
	// 	// get RawBytes from data
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}

	// 	// Now do something with the data.
	// 	// Here we just print each column as a string.
	// 	var value string
	// 	for i, col := range values {
	// 		// Here we can check if the value is nil (NULL value)
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		fmt.Println(columns[i], ": ", value)

	// 	}
	// 	op.Id = string(values[0])
	// 	op.Group_id = string(values[1])
	// 	op.Name = string(values[2])
	// 	op.Path = string(values[3])
	// 	op.Create_time = string(values[4])
	// 	op.Is_delete = string(values[5])

	// 	ops = append(ops, op)

	// 	fmt.Println("-----------------------------------")
	// }

	// fmt.Println(values[0], op)
	// fmt.Println(ops)
}
