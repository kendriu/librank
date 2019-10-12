package main

import (
	"github.com/kendriu/librank/internal/audioteka"
)

type Exporter struct {
	service audioteka.Service
}

func (e *Exporter) Export(exports chan interface{}) {
	e.service.Prune()
	for res := range exports {
		book, ok := res.(*audioteka.Book)
		if ok {
			e.service.Add(*book)
		} else {
			panic(res)
		}
	}
}
