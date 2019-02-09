package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type book struct {
	title    string
	author   []string
	category string
	series   string
}

func getBookParser(c chan book) func(*colly.HTMLElement) {

	parseBook := func(e *colly.HTMLElement) {
		title := e.DOM.Find(".product-title").Text()
		category := e.DOM.Find(".category").Text()
		author := e.DOM.Find("ul.product-spec__data li .text").Slice(1, 2).Text()
		author = strings.TrimSpace(author)
		authors := strings.Split(author, ",")
		for i, v := range authors {
			authors[i] = strings.TrimSpace(v)
		}
		b := book{
			title:    title,
			author:   authors,
			category: category,
		}
	}
	return parseBook

}

func main() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	c := colly.NewCollector(
		colly.AllowedDomains("audioteka.com"),
		colly.Async(true),
		colly.CacheDir("./cache"),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 100,
		// RandomDelay: 5 * time.Second,
	})

	var counter int
	var books []book
	crawl := func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Printf("Found: %q", e.Text)
		e.Request.Visit(link)
	}
	c.OnHTML(".fs17 a", crawl)
	c.OnHTML(".all-categories a", crawl)
	c.OnHTML(".items-loader a", crawl)
	c.OnHTML(".item__title a", crawl)
	c.OnHTML(".view-product", parseBook)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
	c.Visit("https://audioteka.com/pl/audiobooks")
	c.Wait()

	fmt.Printf("Found %v books\n", counter)
}
