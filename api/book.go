package api

import (
	"encoding/xml"
	"fmt"
)

type (
	Book struct {
		ID  string `xml:"id"`
		URL string `xml:"url"`
	}
)

func GetBookURLByISBN(isbn string) (string, error) {
	rawurl := fmt.Sprintf("https://www.goodreads.com/book/isbn/%s?key=%s", isbn, key)
	resp, err := client.Get(rawurl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var grResp Response
	err = xml.NewDecoder(resp.Body).Decode(&grResp)
	if err != nil {
		return "", err
	}

	return grResp.Book.URL, nil
}
