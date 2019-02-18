package api

import (
	"net/http"
	"time"
)

const key = "V6TnFk4YbLS0GdljcCGKQ"

var client = &http.Client{}

func init() {
	client.Timeout = time.Duration(20 * time.Second)
}
