package api

import (
	"encoding/xml"
	"errors"
	"net/http"
	"time"
)

const key = "V6TnFk4YbLS0GdljcCGKQ"

var (
	client = &http.Client{}

	NotFoundError      = errors.New("book not found")
	FailedRequestError = errors.New("request failed")
)

func getBook(rawurl string) (*Book, error) {
	resp, err := client.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return nil, NotFoundError
		}
		return nil, FailedRequestError
	}

	var grResp Response
	err = xml.NewDecoder(resp.Body).Decode(&grResp)
	if err != nil {
		return nil, err
	}

	return &grResp.Book, nil
}

func init() {
	client.Timeout = time.Duration(20 * time.Second)
}
