package coinmarketcap

import "fmt"

func ExampleConvert() {
	client, _ := NewClient()
	defer client.Close()

	data, _ := client.GlobalData(
		Convert("DKK"),
	)

	cap, _ := data.MarketCap("DKK")
	fmt.Printf("Global marketcap: %f\n", cap)
}

func ExampleCurrency() {
	client, _ := NewClient()
	defer client.Close()

	data, _ := client.Ticker(
		Currency("bitcoin"),
	)
	cap, _ := data.Index(0).MarketCap("USD")
	fmt.Printf("Bitcoin Market Cap: %f\n", cap)
}
