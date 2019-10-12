package audioteka

//Book represents a book found by crawler
type Book struct {
	Title    string
	Author   []string
	Category string
	Series   string
	URL      string
}

func NewBook(title string, author []string, category string, url string) *Book {
	return &Book{
		Title:    title,
		Author:   author,
		Category: category,
		URL:      url,
	}
}
