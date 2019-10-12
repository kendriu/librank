package lubimy_czytac

import (
	"sort"
)

type CatalogueItem struct {
	Title   string
	authors []string
}

func (c *CatalogueItem) equal(item CatalogueItem) bool {
	if c.Title != item.Title {
		return false
	}
	if len(c.authors) != len(item.authors) {
		return false
	}
	for _, a := range c.authors {
		for _, a2 := range item.authors {
			if a != a2 {
				return false
			}
		}
	}
	return true
}

func (c *CatalogueItem) addAuthor(author string) {
	c.authors = append(c.authors, author)
	sort.Strings(c.authors)
}

func NewCatalogueItem(title string, authors []string, ) *CatalogueItem {
	item := &CatalogueItem{
		Title: title,
	}
	for _, a := range authors {
		item.addAuthor(a)
	}
	return item
}

type Book struct {
	NeedleTitle string
	Items       []CatalogueItem
}

func (b *Book) Add(item CatalogueItem) {
	b.Items = append(b.Items, item)
}

func (b *Book) updateItems(items []CatalogueItem) {
OUTER:
	for _, item := range items {
		for _, item2 := range b.Items {
			if item.equal(item2) {
				continue OUTER
			}
		}
		b.Items = append(b.Items, item)
	}
}
