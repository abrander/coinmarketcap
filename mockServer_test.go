package coinmarketcap

import (
	"net/http"
	"net/http/httptest"
)

func init() {
	server := httptest.NewServer(http.HandlerFunc(handleMockRequests))

	mockAddress = server.URL
}

var mockAddress string

func handleMockRequests(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {

	case "/ticker/":
		http.ServeFile(w, r, "testdata/ticker-5-EUR.json")
	case "/broken-json/ticker/":
		w.Write([]byte("[}"))

	case "/global/":
		http.ServeFile(w, r, "testdata/global.json")
	case "/broken-json/global/":
		w.Write([]byte("[}"))

	default:
		panic("no handler for " + r.RequestURI)
	}
}
