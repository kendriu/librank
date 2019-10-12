package lubimy_czytac

type Book struct {
	Title  string
	Author []string
}

func NewBook(title string, author []string, ) *Book {
	return &Book{
		Title:  title,
		Author: author,
	}
}
