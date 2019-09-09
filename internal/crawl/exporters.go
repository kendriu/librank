package crawl

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type GOBExporter struct {
	FileName string
}

func (e *GOBExporter) Export(exports chan interface{}) {

	// Create file
	newFile, err := os.OpenFile(e.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("output file creation error: %v", err)
		return
	}

	encoder := gob.NewEncoder(newFile)

	for res := range exports {
		if err := encoder.Encode(res); err != nil {
			log.Printf("gob encoding error on exporter: %v\n", err)
		}
	}
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
