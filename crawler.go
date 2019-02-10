package librank

import (
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/gocolly/colly"
)

// Crawl crawls audioteka for books
func Crawl() Books {
	books := make(Books)
	ch := make(chan Book)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		defer wg.Done()
		for b := range ch {
			if indexed, ok := books[b.URL]; ok {
				log.Warnf("Book already scraped:\nIndexed: %v\nNew:    %v", b, indexed)
				continue
			}
			books[b.URL] = b
		}
	}()
	c := audiotekaCrawler()
	c.OnHTML(".view-product", bookParser(ch))

	c.OnRequest(func(r *colly.Request) {
		log.Debugln("Visiting", r.URL.String())
	})
	if err := c.Visit("https://audioteka.com/pl/audiobooks"); err != nil {
		log.Fatal(err)
	}
	c.Wait()
	close(ch)
	wg.Wait()

	log.Infof("Found %v books\n", len(books))
	return books
}

func audiotekaCrawler() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("audioteka.com"),
		colly.Async(true),
	)
	if CACHE == 1 {
		c.CacheDir = "./cache"
	}
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: PARALLELISM,
	})
	if err != nil {
		log.Fatal(err)
	}

	crawl := func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Debugf("Found: %q", e.Text)
		err := e.Request.Visit(link)
		switch err {
		case nil:
			break
		case colly.ErrAlreadyVisited:
			log.Debugf(err.Error())
		default:
			log.Fatal(err.Error())
		}

	}
	c.OnHTML(".fs17 a", crawl)
	c.OnHTML(".all-categories a", crawl)
	c.OnHTML(".items-loader a", crawl)
	c.OnHTML(".item__title a", crawl)
	return c

}

func bookParser(c chan Book) func(*colly.HTMLElement) {

	parseBook := func(e *colly.HTMLElement) {
		title := e.DOM.Find(".product-title").Text()
		category := e.DOM.Find(".category").Text()
		author := e.DOM.Find("ul.product-spec__data li .text").Slice(1, 2).Text()
		author = strings.TrimSpace(author)
		authors := strings.Split(author, ",")
		for i, v := range authors {
			authors[i] = strings.TrimSpace(v)
		}
		b := Book{
			Title:    title,
			Author:   authors,
			Category: category,
			URL:      URL(e.Request.URL.String()),
		}
		c <- b
	}
	return parseBook
}
