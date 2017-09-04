package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
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
	ExampleScrape()
}
