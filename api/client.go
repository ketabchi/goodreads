package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const key = "V6TnFk4YbLS0GdljcCGKQ"

var (
	Client = &http.Client{}

	NotFoundError      = errors.New("book not found")
	FailedRequestError = errors.New("request failed")
)

func GetDoc(rawurl string) (*goquery.Document, error) {
	res, err := Client.Get(rawurl)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromResponse(res)
}

func getBook(rawurl string) (*Book, error) {
	resp, err := Client.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("%s resulted in %w", rawurl, NotFoundError)
		}
		return nil, fmt.Errorf("%s resulted in %w", rawurl, FailedRequestError)
	}

	var grResp Response
	err = xml.NewDecoder(resp.Body).Decode(&grResp)
	if err != nil {
		return nil, err
	}

	return &grResp.Book, nil
}
