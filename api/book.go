package api

import (
	"encoding/xml"
	"fmt"
)

type (
	Book struct {
		ID    string `xml:"id"`
		URL   string `xml:"url"`
		Title string
		ISBN  string `xml:"isbn"`
		Work
	}

	Work struct {
		BestBookID string `xml:"work>best_book_id"`
	}
)

func GetBookByISBN(isbn string) (Book, error) {
	rawurl := fmt.Sprintf("https://www.goodreads.com/book/isbn/%s?key=%s", isbn, key)
	resp, err := client.Get(rawurl)
	if err != nil {
		return Book{}, err
	}
	defer resp.Body.Close()

	var grResp Response
	err = xml.NewDecoder(resp.Body).Decode(&grResp)
	if err != nil {
		return Book{}, err
	}

	return grResp.Book, nil
}
