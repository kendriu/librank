package main

import (
	"log"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/cache/diskcache"
	"github.com/geziyor/geziyor/export"
	"github.com/kendriu/librank/internal"
	"github.com/kendriu/librank/internal/audioteka"
	"github.com/spf13/viper"
)

func main() {
	internal.Configure()
	log.Print("Scrapping audioteka.pl")

	exporter := &Exporter{audioteka.NewService(
		audioteka.NewScribbleRepository(viper.GetString("TmpPath")))}

	options := &geziyor.Options{
		Cache:                 diskcache.New("./librank_cache"),
		StartURLs:             []string{"http://audioteka.com/pl/audiobooks"},
		ParseFunc:             CategoriesParse,
		Exporters:             []export.Exporter{exporter},
		AllowedDomains:        []string{"audioteka.com"},
		ConcurrentRequests:    viper.GetInt("ConcurrentRequests"),
		LogDisabled:           true,
		CharsetDetectDisabled: true,
	}
	geziyor.NewGeziyor(options).Start()
}
