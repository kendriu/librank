package librank

//URL of a found book
type URL string

//Book represents a book found by crawler
type Book struct {
	Title    string
	Author   []string
	Category string
	Series   string
	URL      URL
}

//Books index of found books
type Books map[URL]Book