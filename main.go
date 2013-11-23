package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

const (
	url = "http://portal.acttv.in/index.php/myusage"
)

func main() {
	log.Printf("Retrieving %v", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tr#total td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		used := s.Text()
		log.Printf("Looks like you have used %v", used)
	})
}
