package api

import "testing"

func TestGetBookURLByISBN(t *testing.T) {
	tests := []struct {
		isbn string
		exp  string
	}{
		{
			"9780132350884",
			"https://www.goodreads.com/book/show/3735293-clean-code",
		},
	}

	for i, test := range tests {
		url, err := GetBookURLByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on getting book url by isbn: %s", i, err)
			continue
		}
		if url != test.exp {
			t.Errorf("Expected url %s, but got %s", test.exp, url)
		}
	}
}
