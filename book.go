package goodreads

import (
	"errors"
	"regexp"

	"github.com/ketabchi/goodreads/api"
	"github.com/ketabchi/util"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	url string
	doc *goquery.Document
	api.Book
}

func NewBookByISBN(isbn string) (*Book, error) {
	gb, err := api.GetBookByISBN(isbn)
	if err != nil {
		return nil, err
	}

	book, err := newBook(gb.URL)
	if err != nil {
		return nil, err
	}
	book.Book = *gb

	return book, nil
}

func NewBookByTitle(title string) (*Book, error) {
	gb, err := api.GetBookByTitle(title)
	if err != nil {
		return nil, err
	}

	book, err := newBook(gb.URL)
	if err != nil {
		return nil, err
	}
	book.Book = *gb

	return book, nil
}

func NewBook(url string) (*Book, error) {
	id := bookID(url)
	if len(id) == 0 {
		return nil, errors.New("can't get book id from url")
	}

	gb, err := api.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	book, err := newBook(gb.URL)
	if err != nil {
		return nil, err
	}
	book.Book = *gb

	return book, nil
}

func newBook(url string) (*Book, error) {
	doc, err := api.GetDoc(url)
	if err != nil {
		return nil, err
	}

	return &Book{url: url, doc: doc}, nil
}

var bookIDRe = regexp.MustCompile("goodreads\\.com\\/book\\/show\\/([0-9]+)")

func bookID(url string) string {
	ss := bookIDRe.FindStringSubmatch(url)
	if len(ss) < 2 {
		return ""
	}
	return ss[1]
}

func (b *Book) Genres() []string {
	genres := make([]string, 0)

	b.doc.Find(".left .bookPageGenreLink").Each(func(i int, sel *goquery.Selection) {
		if g := sel.Text(); !util.ExistsStringInSlice(g, genres) {
			genres = append(genres, g)
		}
	})

	return genres
}

func (b *Book) HasAuthor(name string) bool {
	authors := b.Authors
	for _, author := range authors {
		if author.Name == name {
			return true
		}
	}
	return false
}
