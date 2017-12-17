package coinmarketcap

import (
	"fmt"
)

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
