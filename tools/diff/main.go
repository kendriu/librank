package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"librank"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func loadBooks(path string) librank.Books {
	var read []byte
	var err error
	if read, err = ioutil.ReadFile(path); err != nil {
		log.Panic(err)
	}
	buf := bytes.NewBuffer(read)
	dec := gob.NewDecoder(buf)
	decoded := make(librank.Books)
	if err := dec.Decode(&decoded); err != nil {
		log.Panic(err)
	}
	return decoded

}

func main() {
	librank.SetupLogger()
	path := os.Args[1]

	series := make([]librank.Books, 10)
	for i := 0; i <= 9; i++ {
		series[i] = loadBooks(path + "_" + strconv.Itoa(i+1))
	}
	for _, books := range series {
		fmt.Println(len(books))
	}
}
