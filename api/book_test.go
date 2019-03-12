package api

import (
	"testing"
)

func TestGetBookByISBN(t *testing.T) {
	tests := []struct {
		isbn   string
		expURL string
	}{
		{
			"9780132350884",
			"https://www.goodreads.com/book/show/3735293-clean-code",
		},
	}

	for i, test := range tests {
		book, err := GetBookByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on getting book url by isbn: %s", i, err)
			continue
		}
		if url := book.URL; url != test.expURL {
			t.Errorf("Expected url %s, but got %s", test.expURL, url)
		}
	}
}

func TestGetBookByISBNError(t *testing.T) {
	isbn := "9789643416100"
	_, err := GetBookByISBN(isbn)
	if err != BookNotFoundError {
		t.Errorf("Expected book not found error but got: %s", err)
	}
}
