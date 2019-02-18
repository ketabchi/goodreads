package goodreads

import (
	"testing"

	"github.com/ketabchi/util"
)

func TestGenres(t *testing.T) {
	tests := []struct {
		url string
		exp []string
	}{
		{
			"https://www.goodreads.com/book/show/7683075",
			[]string{"Philosophy", "Nonfiction"},
		},
		{
			"https://www.goodreads.com/book/show/178493._",
			[]string{"Novels"},
		},
		{
			"https://www.goodreads.com/book/show/10541690",
			[]string{"Classics", "Fiction", "Academic", "School",
				"Literature", "Historical", "Historical Fiction",
				"Read For School", "Novels", "American",
				"Young Adult", "High School", "Classic Literature"},
		},
	}

	for i, test := range tests {
		b, err := newBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book: %s", i, err)
			continue
		}
		if genres := b.Genres(); !util.CheckSliceEq(genres, test.exp) {
			t.Errorf("Test %d: Expected genres %q, but got %q", i, test.exp, genres)
		}
	}
}

func TestBestBookID(t *testing.T) {
	tests := []struct {
		isbn string
		exp  string
	}{
		{
			"9780062515674",
			"43877",
		},
	}

	for i, test := range tests {
		book, err := NewBookByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on creating book: %s", i, err)
			continue
		}

		if book.Work.BestBookID != test.exp {
			t.Errorf("Test %d: Expected book best book id %s, but got %s",
				i, test.exp, book.Work.BestBookID)
		}
	}
}
