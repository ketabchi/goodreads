package api

import "os"

var (
	proxy_addr = ""
	username   = ""
	password   = ""
)

func init() {
	if proxy_addr == "" {
		proxy_addr = os.Getenv("PROXY_ADDR")
		username = os.Getenv("PROXY_USERNAME")
		password = os.Getenv("PROXY_PASSWORD")
	}
}
