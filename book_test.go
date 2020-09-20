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
			[]string{"Philosophy", "Nonfiction", "Unfinished",
				"Psychology", "Classics"},
		},
		{
			"https://www.goodreads.com/book/show/178493._",
			[]string{"Novels", "Fiction"},
		},
		{
			"https://www.goodreads.com/book/show/10541690",
			[]string{"Classics", "Fiction", "Academic", "School",
				"Literature", "Historical", "Historical Fiction"},
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

func TestNewBook(t *testing.T) {
	tests := []struct {
		url           string
		expBestBookID string
	}{
		{
			"https://www.goodreads.com/book/show/19060872-the-monk-who-sold-his-ferrari",
			"43877",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book: %s", i, err)
			continue
		}

		if book.Work.BestBookID != test.expBestBookID {
			t.Errorf("Test %d: Expected book best book id %s, but got %s",
				i, test.expBestBookID, book.Work.BestBookID)
		}
	}
}

func TestBookID(t *testing.T) {
	tests := []struct {
		url string
		exp string
	}{
		{
			"https://www.goodreads.com/book/show/38750904-the-accidental-further-adventures-of-the-hundred-year-old-man",
			"38750904",
		},
		{
			"https://www.goodreads.com/book/show/1.Harry_Potter_and_the_Half_Blood_Prince",
			"1",
		},
		{
			"https://www.goodreads.com/book/show/3",
			"3",
		},
	}

	for i, test := range tests {
		id := bookID(test.url)
		if id != test.exp {
			t.Errorf("Test %d: Expected id %s, but got %s", i, test.exp, id)
		}
	}
}

func TestHasAuthor(t *testing.T) {
	tests := []struct {
		url    string
		author string
		exp    bool
	}{
		{
			"https://www.goodreads.com/book/show/38750904-the-accidental-further-adventures-of-the-hundred-year-old-man",
			"Jonas Jonasson",
			true,
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Couldnt create book from %s: %s", i, test.url, err)
			continue
		}

		if exists := book.HasAuthor(test.author); exists != test.exp {
			t.Errorf("Test %d: Expected has author %v, but got %v", i, test.exp, exists)
		}
	}
}

func TestSerie(t *testing.T) {
	tests := []struct {
		url string
		exp string
	}{
		{
			"https://www.goodreads.com/book/show/17347634-me-before-you",
			"Me Before You",
		},
		{
			"https://www.goodreads.com/book/show/2195738._04_",
			"Tintin",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Couldnt create book from %s: %s", i, test.url, err)
			continue
		}

		if serie := book.Serie(); serie != test.exp {
			t.Errorf("Test %d: Expected serie %s, but got %s", i, test.exp, serie)
		}
	}
}
