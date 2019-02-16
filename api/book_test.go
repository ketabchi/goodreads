package api

import "testing"

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
		url, err := GetBookByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on getting book url by isbn: %s", i, err)
			continue
		}
		if url != test.expURL {
			t.Errorf("Expected url %s, but got %s", test.expURL, url)
		}
	}
}
