package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/kendriu/librank/internal/audioteka"
)

func CategoriesParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".view-product").Each(func(_ int, s *goquery.Selection) {
		aBook := getBook(s, r.Request.URL.String())
		for _, author := range aBook.Author {
			switch author {
			case "CIDEB EDITRICE",
				"Gamma",
				"Disney":
				return
			}
		}
		g.Exports <- aBook
	})

	selector := ".fs17 a, .all-categories a, .items-loader a, .item__title a"
	r.HTMLDoc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			switch {
			case strings.Contains(href, "angielski-n"),
				strings.Contains(href, "chinski-n"),
				strings.Contains(href, "dla-dzieci"),
				strings.Contains(href, "do-rzeczy"),
				strings.Contains(href, "hiszanski-n"),
				strings.Contains(href, "inne-jezyki-n"),
				strings.Contains(href, "ksiazki-anglojezyczne"),
				strings.Contains(href, "nauka-jezykow-n"),
				strings.Contains(href, "niemecki-n"),
				strings.Contains(href, "polityka-n"),
				strings.Contains(href, "prasa-n"),
				strings.Contains(href, "rosyjski-n"),
				strings.Contains(href, "wiek-powyzej-lat-7"),
				strings.Contains(href, "wierszyki-i-zarciki"),
				strings.Contains(href, "wloski-n"),
				strings.Contains(href, "wprost-n"):
				return
			default:
				g.Get(href, CategoriesParse)
			}
		}
	})
}

func getBook(s *goquery.Selection, URL string) *audioteka.Book {
	author := s.Find("ul.product-spec__data li .text").Slice(1, 2).Text()
	author = strings.TrimSpace(author)
	authors := strings.Split(author, ",")
	for i, v := range authors {
		authors[i] = strings.TrimSpace(v)
	}
	//series := s.Find("ul.product-spec__data li .text").Slice(5, 6).Text()
	//series = strings.TrimSpace(series)
	log.Println("Found book")

	return audioteka.NewBook(
		s.Find(".product-title").Text(),
		authors,
		s.Find(".category").Text(),
		URL,
	)
}
