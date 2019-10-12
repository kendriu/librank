package lubimy_czytac

import (
	"os"
	"path/filepath"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Repository interface {
	prune()
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

func (s ScribbleRepository) prune() {
	dbPath := filepath.Join(s.dir, s.collectionName)
	if err := os.RemoveAll(dbPath); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

}
