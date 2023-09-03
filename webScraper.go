package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type News struct {
	Link string
	Text string
}

func main() {
	url := "https://www.ambebi.ge/"
	keyWord := " "

	c := colly.NewCollector()

	var newsArr []News

	c.OnHTML("a", func(e *colly.HTMLElement) {
		var news News
		news.Link = e.Attr("href")

		e.DOM.Find("p").Each(func(index int, item *goquery.Selection) {
			news.Text = strings.TrimSpace(item.Text())
		})

		if news.Text != "" {
			newsArr = append(newsArr, news)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal("Error visiting site:", err)
	}

	fmt.Println("Number of News:", len(newsArr))

	for _, news := range newsArr {
		if strings.Contains(news.Text, keyWord) {
			fmt.Println(news)
		}
	}
}
