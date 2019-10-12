package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/kendriu/librank/internal/audioteka"
	"github.com/kendriu/librank/internal/lubimy_czytac"
)

type RequestsProvider struct {
	audiotekaService audioteka.Service
}

func (r RequestsProvider) startRequests(g *geziyor.Geziyor) {
	titles := r.audiotekaService.GetTitles()
	titles = titles[1:3]
	for _, t := range titles {
		parser := FoundParser{needleTitle: t}
		g.Get(fmt.Sprintf("http://lubimyczytac.pl/szukaj/ksiazki/1?phrase=%s", t), parser.parse)
	}
}

type FoundParser struct {
	needleTitle string
	foundBooks  []*lubimy_czytac.Book
}

func (f *FoundParser) parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".book-general-data").Each(func(_ int, s *goquery.Selection) {
		allAuthors := s.Find(".author")
		var authors []string
		allAuthors.Find("a").Each(func(_ int, s *goquery.Selection) {
			authors = append(authors, s.Text())
		})

		book := lubimy_czytac.NewBook(
			s.Find(".bookTitle").Text(),
			authors,
		)
		f.foundBooks = append(f.foundBooks, book)
		g.Exports <- f
	})
}
