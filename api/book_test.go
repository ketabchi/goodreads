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
	if err != NotFoundError {
		t.Errorf("Expected book not found error but got: %s", err)
	}
}

func TestGetBookByTitle(t *testing.T) {
	tests := []struct {
		title  string
		expURL string
	}{
		{
			"The Accidental Further Adventures of The Hundred Year Old Man",
			"https://www.goodreads.com/book/show/38750904-the-accidental-further-adventures-of-the-hundred-year-old-man",
		},
		{
			"Artemis fowl: the lost colony",
			"https://www.goodreads.com/book/show/37586.The_Lost_Colony",
		},
	}

	for i, test := range tests {
		b, err := GetBookByTitle(test.title)
		if err != nil {
			t.Errorf("Test %d: Error on getting book by title %s: %s",
				i, test.title, err)
			continue
		}

		if url := b.URL; url != test.expURL {
			t.Errorf("Test %d: Expected book url %s, but got %s", i, test.expURL, url)
		}
	}
}
