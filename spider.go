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

func main() {
	// ExampleScrape()

	db, err := sql.Open("mysql", "admin:Dream1tPossible@tcp(114.215.154.110:3306)/first-go")
	rows, err := db.Query("select * from one_piece_pic where id > ?", 1)

	s, err := rows.Columns()
	for i := 0; i < 7; i++ {
		rows.Scan(s)
		rows.Next()

	}
	fmt.Println(rows, err)
}
