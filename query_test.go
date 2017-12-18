package coinmarketcap

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	cases := []struct {
		options  []func(*query)
		expected string
	}{
		{[]func(*query){}, ""},
		{[]func(*query){Start(50)}, "?start=50"},
		{[]func(*query){Start(50), Limit(12)}, "?limit=12&start=50"},
		{[]func(*query){Convert("USD")}, "?convert=USD"},
		{[]func(*query){Currency("bitcoin")}, "bitcoin"},
		{[]func(*query){Currency("bitcoin"), Start(50), Limit(12), Convert("DKK")}, "bitcoin?convert=DKK&limit=12&start=50"},
	}

	for i, c := range cases {
		q := newQuery(c.options)
		if q.tickerQuery() != c.expected {
			t.Errorf("%d Got wrong query, got '%s', expected '%s'", i, q.tickerQuery(), c.expected)
		}
	}
}

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
