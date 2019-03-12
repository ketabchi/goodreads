package api

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type (
	Book struct {
		ID    string `xml:"id"`
		URL   string `xml:"url"`
		Title string `xml:"title"`
		ISBN  string `xml:"isbn13"`
		Work
	}

	Work struct {
		BestBookID  string `xml:"work>best_book_id"`
		RatingCount int    `xml:"work>ratings_count"`
	}
)

var (
	BookNotFoundError = errors.New("book not found")
)

func GetBookByISBN(isbn string) (*Book, error) {
	rawurl := fmt.Sprintf("https://www.goodreads.com/book/isbn/%s?key=%s", isbn, key)
	resp, err := client.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, BookNotFoundError
	}

	var grResp Response
	err = xml.NewDecoder(resp.Body).Decode(&grResp)
	if err != nil {
		return nil, err
	}

	return &grResp.Book, nil
}
