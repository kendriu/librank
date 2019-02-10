package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"librank"
)

func main() {

	buf := new(bytes.Buffer)

	dec := gob.NewDecoder(buf)
	var decoded map[string]librank.Book
	if err := dec.Decode(&decoded); err != nil {
		panic(err)
	}
	fmt.Println(decoded)
}
