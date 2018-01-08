package coinmarketcap

import (
	"fmt"
	"testing"
	"time"
)

func TestBaseURL(t *testing.T) {
	client, err := NewClient(BaseURL("new-url"))
	if err != nil {
		t.Fatalf("NewClient() returned an error: %s", err.Error())
	}
	defer client.Close()

	if client.baseURL != "new-url" {
		t.Errorf("Failed to set BaseURL, fgot %s, expected %s", client.baseURL, "new-url")
	}
}

func TestRateLimit(t *testing.T) {
	client, _ := NewClient(BaseURL(mockAddress), RateLimit(10000))
	defer client.Close()

	client.Ticker()

	// Allow throttler to tick once
	time.Sleep(time.Millisecond * 100)
}

func TestTicker(t *testing.T) {
	client, _ := NewClient(BaseURL(mockAddress))
	defer client.Close()

	coins, err := client.Ticker()
	if err != nil {
		t.Fatalf("Ticker() returned an unexpected error: %s", err.Error())
	}

	if len(coins) != 5 {
		t.Fatalf("Ticker() not not return the expected amount of coins, got %d, expected 5", len(coins))
	}
}

func TestTickerBrokenJSON(t *testing.T) {
	client, _ := NewClient(BaseURL(mockAddress + "/broken-json"))
	defer client.Close()

	_, err := client.Ticker()
	if err == nil {
		t.Errorf("Ticker() did not fail as expected")
	}
}

func TestTickerFail(t *testing.T) {
	client, _ := NewClient(BaseURL("http://127.0.0.1:0/"))
	defer client.Close()

	_, err := client.Ticker()
	if err == nil {
		t.Errorf("Ticker() did not fail as expected")
	}
}

func TestGlobalData(t *testing.T) {
	client, _ := NewClient(BaseURL(mockAddress))
	defer client.Close()

	_, err := client.GlobalData()
	if err != nil {
		t.Fatalf("GlobalData() returned an unexpected error: %s", err.Error())
	}
}

func TestGlobalDataBrokenJSON(t *testing.T) {
	client, _ := NewClient(BaseURL(mockAddress + "/broken-json"))
	defer client.Close()

	_, err := client.GlobalData()
	if err == nil {
		t.Errorf("GlobalData() did not fail as expected")
	}
}

func TestGlobalDataFail(t *testing.T) {
	client, _ := NewClient(BaseURL("http://127.0.0.1:0/"))
	defer client.Close()

	_, err := client.GlobalData()
	if err == nil {
		t.Errorf("GlobalData() did not fail as expected")
	}
}

func TestGraph(t *testing.T) {
	client, _ := NewClient(GraphBaseURL(mockAddress))
	defer client.Close()

	graph, err := client.Graph("bitcoin", time.Time{}, time.Time{})
	if err != nil {
		t.Fatalf("Graph() returned an unexpected error: %s", err.Error())
	}

	if len(graph.MarkepCap) != 9 {
		t.Fatalf("Graph() did not return the expected amount of points, got %d, expected 9", len(graph.MarkepCap))
	}
}

func TestGraphBrokenJSON(t *testing.T) {
	client, _ := NewClient(GraphBaseURL(mockAddress + "/broken-json"))
	defer client.Close()

	_, err := client.Graph("bitcoin", time.Time{}, time.Time{})
	if err == nil {
		t.Errorf("Graph() did not fail as expected")
	}
}

func TestGraphFail(t *testing.T) {
	client, _ := NewClient(GraphBaseURL("http://127.0.0.1:0/"))
	defer client.Close()

	_, err := client.Graph("bitcoin", time.Time{}, time.Time{})
	if err == nil {
		t.Errorf("Graph() did not fail as expected")
	}
}

func ExampleNewClient() {
	client, _ := NewClient()
	defer client.Close()
}

func ExampleClient_Ticker() {
	client, _ := NewClient()
	defer client.Close()

	// Per API documentation ticker will return 100 results per default.
	top100, _ := client.Ticker()
	for _, coin := range top100 {
		fmt.Printf("%d: %s [%s]\n", coin.Rank, coin.Symbol, coin.ID)
	}
}

func ExampleClient_Ticker_top5() {
	client, _ := NewClient()
	defer client.Close()

	top, _ := client.Ticker(
		Limit(5),
	)
	for _, coin := range top {
		fmt.Printf("%d: %s [%s]\n", coin.Rank, coin.Symbol, coin.ID)
	}
}

func ExampleClient_Ticker_kroner() {
	client, _ := NewClient()
	defer client.Close()

	result, _ := client.Ticker(
		Currency("bitcoin"), // We're only interested in bitcoin.
		Convert("DKK"),      // Lets have output in danish kroner.
	)
	coin := result.Index(0)
	price, _ := coin.Price("DKK")
	fmt.Printf("Current price in DKK: %f\n", price)
}

func ExampleClient_GlobalData() {
	client, _ := NewClient()
	defer client.Close()

	data, _ := client.GlobalData()
	fmt.Printf("CoinMarketCap currently lists %d assets\n", data.ActiveAssets)
}
