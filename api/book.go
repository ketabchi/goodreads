package api

import (
	"errors"
	"fmt"
	"net/url"
)

type (
	Book struct {
		ID      string   `xml:"id"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		ISBN    string   `xml:"isbn13"`
		Work    Work     `xml:"work"`
		Authors []Author `xml:"authors>author"`
	}

	Work struct {
		BestBookID  string `xml:"best_book_id"`
		RatingCount int    `xml:"ratings_count"`
	}

	Author struct {
		Name string `xml:"name"`
		Role string `xml:"role"`
	}
)

func GetBookByISBN(isbn string) (*Book, error) {
	rawurl := fmt.Sprintf("https://www.goodreads.com/book/isbn/%s?key=%s", isbn, key)

	return getBook(rawurl)
}

func GetBookByTitle(args ...string) (*Book, error) {
	if len(args) == 0 {
		return nil, errors.New("get book by title requires at least title")
	}
	title := url.QueryEscape(args[0])
	author := ""
	if len(args) > 1 {
		author = url.QueryEscape(args[1])
	}

	rawurl := fmt.Sprintf("https://www.goodreads.com/book/title.xml?title=%s&author=%s&key=%s", title, author, key)

	return getBook(rawurl)
}

func GetBookByID(id string) (*Book, error) {
	rawurl := fmt.Sprintf("https://www.goodreads.com/book/show/%s.xml?key=%s", id, key)

	return getBook(rawurl)
}
