package helper

import (
	"net/http"
	"time"
)

var HttpClient *http.Client

func HttpInit() {
	HttpClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
}
