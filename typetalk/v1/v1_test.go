package v1

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const fixturesPath = "../../testdata/v1/"

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	client.SetTypetalkToken("DUMMY_TOKEN")
	parsedURL, _ := url.Parse(server.URL)
	client.client.BaseURL = parsedURL
}

func teardown() {
	server.Close()
}
