package internal

import (
	"github.com/sirupsen/logrus"
)

type GOBExporter struct {
	FileName string
}

type CountingExporter struct {
	Counter int
}

func (e *CountingExporter) Export(exports chan interface{}) {
	for range exports {
		e.Counter = e.Counter + 1
	}
}

type LoggingExporter struct{}

func (e *LoggingExporter) Export(exports chan interface{}) {
	for res := range exports {
		logrus.Print(res)
	}
}
