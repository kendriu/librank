package main

import (
	"github.com/fpfeng/httpcache/diskcache"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/export"
	"github.com/kendriu/librank/internal/audioteka"
	"github.com/kendriu/librank/internal/crawl"
)


func main() {
	//exporter := librank.GOBExporter{FileName: "./tmp/audioteka.gob"}

	exporter := &export.JSON{FileName:"./tmp/audioteka.json"}
	//exporter := &export.PrettyPrint{}
	geziyor.NewGeziyor(&geziyor.Options{
		Cache:                 diskcache.New("./librank_cache"),
		StartURLs:             []string{"http://audioteka.com/pl/audiobooks"},
		ParseFunc:             audioteka.CategoriesParse,
		Exporters:             []export.Exporter{exporter},
		AllowedDomains:        []string{"audioteka.com"},
		ConcurrentRequests:    crawl.PARALLELISM,
		LogDisabled:           true,
		CharsetDetectDisabled: true,
	}).Start()
}
