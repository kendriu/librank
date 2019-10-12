package audioteka

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Repository interface {
	save(book Book)
	all() []Book
	prune()
}

type JsonRepository struct {
	path string
}

func NewJsonRepository(path string) *JsonRepository {
	return &JsonRepository{path: path}
}

func (j JsonRepository) save(book Book) {

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString("text to append\n"); err != nil {
		panic(err)
	}
}

func (j JsonRepository) all() []Book {
	panic("implement me")
}

func (j JsonRepository) prune() {
	if err := os.Remove(j.path); ! os.IsNotExist(err) {
		panic(err)
	}
}

type ScribbleRepository struct {
	dir            string
	db             *scribble.Driver
	collectionName string
}

func (s ScribbleRepository) save(book Book) {
	resource := s.toResource(book)
	if err := s.db.Write(s.collectionName, resource, book); err != nil {
		panic(err)
	}
}

func (s ScribbleRepository) toResource(book Book) string {
	resource := book.Title
	resource = strings.ReplaceAll(resource, "/", "_")
	return resource
}

func (s ScribbleRepository) all() []Book {
	records, err := s.db.ReadAll(s.collectionName)
	if err != nil {
		panic(err)
	}

	var books []Book
	for _, f := range records {
		bookFound := Book{}
		if err := json.Unmarshal([]byte(f), &bookFound); err != nil {
			panic(err)
		}
		books = append(books, bookFound)
	}
	return books
}

func (s ScribbleRepository) prune() {
	dbPath := filepath.Join(s.dir, s.collectionName)
	if err := os.RemoveAll(dbPath); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

}

func NewScribbleRepository(dir string) *ScribbleRepository {
	db, err := scribble.New(dir, &scribble.Options{})
	if err != nil {
		panic(err)
	}
	return &ScribbleRepository{dir: dir, db: db, collectionName: "audioteka"}
}
