package goodreads

import (
	"testing"
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
		b, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book: %s", i, err)
			continue
		}
		if genres := b.Genres(); !checkEq(genres, test.exp) {
			t.Errorf("Test %d: Expected genres %q, but got %q", i, test.exp, genres)
		}
	}
}

func checkEq(ss1, ss2 []string) bool {
	if len(ss1) != len(ss2) {
		return false
	}

	for _, s1 := range ss1 {
		found := false
		for _, s2 := range ss2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
