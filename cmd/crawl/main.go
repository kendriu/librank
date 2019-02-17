package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"librank"
	"os"
	"strconv"
)

func main() {
	librank.SetupLogger()
	path := os.Args[1]

	for i := 1; i <= 10; i++ {
		books := librank.Crawl()
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(books)
		ioutil.WriteFile(path+"_"+strconv.Itoa(i), buf.Bytes(), 0600)
	}
}
