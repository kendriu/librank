package audioteka

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/kendriu/librank/internal/book"
)

func CategoriesParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".view-product").Each(func(i int, s *goquery.Selection) {
		aBook := BookParse(s)
		for _, author := range aBook.Author {
			switch author {
			case "CIDEB EDITRICE",
				"Gamma",
				"Disney":
				return
			}
		}
		aBook.URL = book.URL(r.Request.URL.String())
		g.Exports <- aBook
	})

	selector := ".fs17 a, .all-categories a, .items-loader a, .item__title a"
	r.HTMLDoc.Find(selector).Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			switch {
			case strings.Contains(href, "prasa-n"),
				strings.Contains(href, "polityka-n"),
				strings.Contains(href, "do-rzeczy"),
				strings.Contains(href, "wprost-n"):
				return
			default:
				g.Get(href, CategoriesParse)
			}
		}
	})
}

func BookParse(s *goquery.Selection) book.Book {

	author := s.Find("ul.product-spec__data li .text").Slice(1, 2).Text()
	author = strings.TrimSpace(author)
	authors := strings.Split(author, ",")
	for i, v := range authors {
		authors[i] = strings.TrimSpace(v)
	}
	//series := s.Find("ul.product-spec__data li .text").Slice(5, 6).Text()
	//series = strings.TrimSpace(series)

	return book.Book{
		Title:    s.Find(".product-title").Text(),
		Author:   authors,
		Category: s.Find(".category").Text(),
		//Series:   series,
	}
}
