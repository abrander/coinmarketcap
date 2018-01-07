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

	case "/currencies/bitcoin/-62135596800000/-62135596800000":
		http.ServeFile(w, r, "testdata/bitcoin-graph.json")
	case "/broken-json/currencies/bitcoin/-62135596800000/-62135596800000":
		w.Write([]byte("[}"))

	default:
		panic("no handler for " + r.RequestURI)
	}
}
