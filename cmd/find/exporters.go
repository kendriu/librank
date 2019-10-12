package main

import (
	"github.com/kendriu/librank/internal/lubimy_czytac"
)

type Exporter struct {
	service lubimy_czytac.Service
}

func (e *Exporter) Export(exports chan interface{}) {
	e.service.Prune()
	for res := range exports {
		book, ok := res.(*lubimy_czytac.Book)
		if ok {
			e.service.Update(book)
		} else {
			panic(res)
		}
	}
}
