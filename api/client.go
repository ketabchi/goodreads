package api

import (
	"encoding/xml"
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

const key = "V6TnFk4YbLS0GdljcCGKQ"

var (
	client = &http.Client{Timeout: time.Second * 20}

	NotFoundError      = errors.New("book not found")
	FailedRequestError = errors.New("request failed")
)

func GetDoc(rawurl string) (*goquery.Document, error) {
	res, err := client.Get(rawurl)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromResponse(res)
}

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

func addProxy() {
	auth := &proxy.Auth{username, password}
	dial, err := proxy.SOCKS5("tcp", proxy_addr, auth, proxy.Direct)
	if err != nil {
		log.Fatalf("Error on creating proxy dialer: %s", err)
	}

	httpTransport := &http.Transport{}
	httpTransport.Dial = dial.Dial

	client.Transport = httpTransport
}

func init() {
	if proxy_addr != "" {
		addProxy()
	}
}
