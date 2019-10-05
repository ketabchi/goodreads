package api

import "os"

var (
	proxy_addr = os.Getenv("PROXY_ADDR")
	username   = os.Getenv("PROXY_USERNAME")
	password   = os.Getenv("PROXY_PASSWORD")
)
