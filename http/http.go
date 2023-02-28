package http

import (
	"net/http"
	"strings"
)

var HttpClientInstance = &http.Client{}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func Get(url string, headers map[string]string) (*http.Response, error) {
	return DoRequest(HttpClientInstance, "GET", url, headers, "")
}

func Post(url string, headers map[string]string, body string) (*http.Response, error) {
	return DoRequest(HttpClientInstance, "POST", url, headers, body)
}

func DoRequest(httpClient HttpClient, method string, url string, headers map[string]string, body string) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)

	return resp, err
}
