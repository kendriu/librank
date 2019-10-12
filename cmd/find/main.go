package main

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/cache/diskcache"
	"github.com/geziyor/geziyor/export"
	"github.com/kendriu/librank/internal"
	"github.com/kendriu/librank/internal/lubimy_czytac"
	"github.com/spf13/viper"
)

func main() {
	internal.Configure()
	exporter := &export.JSON{FileName: viper.GetString("FoundBooksPath")}
	options := &geziyor.Options{
		Cache:             diskcache.New("./librank_cache"),
		StartRequestsFunc: lubimy_czytac.StartRequests,
		//ParseFunc:             audioteka.CategoriesParse,
		Exporters:             []export.Exporter{exporter},
		AllowedDomains:        []string{"lubimyczytac.pl"},
		ConcurrentRequests:    1,
		LogDisabled:           true,
		CharsetDetectDisabled: true,
	}
	g := geziyor.NewGeziyor(options)
	g.Start()

}
