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
