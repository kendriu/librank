package lubimy_czytac

import (
	"os"
	"path/filepath"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Repository interface {
	prune()
	get(needleTitle string) *Book
	save(book *Book)
}

type ScribbleRepository struct {
	dir            string
	db             *scribble.Driver
	collectionName string
}

func NewScribbleRepository(dir string) *ScribbleRepository {
	db, err := scribble.New(dir, &scribble.Options{})
	if err != nil {
		panic(err)
	}
	return &ScribbleRepository{dir: dir, db: db, collectionName: "lubimy_czytac"}
}

func (s *ScribbleRepository) get(title string) *Book {
	var book Book
	if err := s.db.Read(s.collectionName, title, &book); err != nil {
		book.NeedleTitle = title
	}
	return &book
}

func (s *ScribbleRepository) save(book *Book) {
	resource := s.toResource(*book)
	if err := s.db.Write(s.collectionName, resource, book); err != nil {
		panic(err)
	}
}

func (s ScribbleRepository) toResource(book Book) string {
	resource := book.NeedleTitle
	resource = strings.ReplaceAll(resource, "/", "_")
	return resource
}


func (s *ScribbleRepository) prune() {
	dbPath := filepath.Join(s.dir, s.collectionName)
	if err := os.RemoveAll(dbPath); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

}
