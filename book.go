package goodreads

import (
	"github.com/ketabchi/goodreads/api"

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

// Not exporting this because when calling newBook directly book.Book doesn't
// get filled
func newBook(url string) (*Book, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	return &Book{url: url, doc: doc}, nil
}

func (b *Book) Genres() []string {
	genres := make([]string, 0)

	b.doc.Find(".left .bookPageGenreLink").Each(func(i int, sel *goquery.Selection) {
		if g := sel.Text(); !exists(g, genres) {
			genres = append(genres, g)
		}
	})

	return genres
}

func exists(s string, ss []string) bool {
	for _, s1 := range ss {
		if s1 == s {
			return true
		}
	}
	return false
}
