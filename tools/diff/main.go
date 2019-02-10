package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"librank"
	"os"
	"sort"
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
func printMissing(a, b librank.Books) {
	for url := range a {
		if _, ok := b[url]; !ok {
			fmt.Println("\t" + a[url].Title)
		}
	}
}
func printDiff(a, b librank.Books) {
	fmt.Println("Missing in first:")
	printMissing(b, a)

	fmt.Println("Missing in second:")
	printMissing(a, b)
}

func main() {
	librank.SetupLogger()
	path := os.Args[1]

	series := make([]librank.Books, 10)
	for i := 0; i <= 9; i++ {
		series[i] = loadBooks(path + "_" + strconv.Itoa(i+1))
	}
	sort.Slice(series, func(i, j int) bool {
		return len(series[i]) < len(series[j])
	})
	var sample librank.Books
	sample = series[8]
	for _, books := range series {
		printDiff(sample, books)
	}
}
