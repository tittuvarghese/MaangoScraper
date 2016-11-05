package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func MaangoScrape(ScrapyURL string, PageCount int) {
	doc, err := goquery.NewDocument(ScrapyURL)
	if err != nil {
		log.Fatal(err)
	}

	// Find the Music collections
	doc.Find("#dle-content .post-serial .img-serial .left-btn-play").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		MovieName := s.Find("a").Text()
		MovieSingleURL, ok := s.Find("a").Attr("href")
		if ok {
			DownloadURL := GetDownloadURL(MovieSingleURL)
			// If the download link not available the maango.me site generates a static ad URL, we are excluding it.
			if DownloadURL != "http://newtemplates.ru/" {
				fmt.Printf("%d. %s - %s\n", PageCount+i+1, MovieName, DownloadURL)
				// write the whole body at once
				f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					panic(err)
				}

				defer f.Close()
				if _, err = f.WriteString(DownloadURL); err != nil {
					panic(err)
				}
				if _, err = f.WriteString("\n"); err != nil {
					panic(err)
				}
			}
		}
	})
}

func GetDownloadURL(SingleItemURL string) string {
	var DownloadURL string
	doc, err := goquery.NewDocument(SingleItemURL)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".moredl").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the download url
		DownloadURL, _ = s.Find("a").Attr("href")
	})
	return DownloadURLF
}

func main() {
	var BaseUrl = "http://www.maango.me"
	var Language = "/hindi" // Set the language http://www.maango.me/hindi/
	var Pages = "/page/"
	// Set the PageCount limit to total available pages.
	for PageCount := 1; PageCount <= 10; PageCount++ {
		var QueryURL = BaseUrl + Language + Pages + strconv.Itoa(PageCount) + "/"
		MaangoScrape(QueryURL, ((PageCount - 1) * 12))
	}
}
