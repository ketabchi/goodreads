package api

import "os"

const (
	proxy_addr = "ad.socadd.com:443"
	username   = "ir435166"
	password   = "529233"
)

func init() {
	if proxy_addr == "" {
		proxy_addr = os.Getenv("PROXY_ADDR")
		username = os.Getenv("PROXY_USERNAME")
		password = os.Getenv("PROXY_PASSWORD")
	}
}
