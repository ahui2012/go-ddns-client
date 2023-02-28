package pubip

import (
	"io"

	"github.com/ahui2012/go-ddns-client/http"
)

var RequestHeaders = map[string]string{
	"Accept":          "text/html",
	"Accept-Charset":  "UTF-8",
	"Accept-Language": "en-US,en",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6)",
}

func GetPublicIP(urls []string) string {
	for _, url := range urls {
		resp, err := http.Get(url, RequestHeaders)
		if err == nil {
			defer resp.Body.Close()
			body, e := io.ReadAll(resp.Body)
			if e == nil {
				return string(body)
			}
		}
	}
	return ""
}
