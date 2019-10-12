package main

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/cache/diskcache"
	"github.com/geziyor/geziyor/export"
	"github.com/kendriu/librank/internal"
	"github.com/kendriu/librank/internal/audioteka"
	"github.com/kendriu/librank/internal/lubimy_czytac"
	"github.com/spf13/viper"
)

func main() {
	internal.Configure()
	tmpPath :=viper.GetString("TmpPath")
	requestsProvider := RequestsProvider{
		audioteka.NewService(
			audioteka.NewScribbleRepository(tmpPath))}

	exporter := &Exporter{lubimy_czytac.NewService(
		lubimy_czytac.NewScribbleRepository(tmpPath))}

	options := &geziyor.Options{
		Cache:                 diskcache.New("./tmp/librank_cache"),
		StartRequestsFunc:     requestsProvider.startRequests,
		Exporters:             []export.Exporter{exporter},
		AllowedDomains:        []string{"lubimyczytac.pl"},
		ConcurrentRequests:    1,
		LogDisabled:           true,
		CharsetDetectDisabled: true,
	}
	g := geziyor.NewGeziyor(options)
	g.Start()

}
