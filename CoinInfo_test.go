package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestCoinInfoAccessors(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ticker-1-EUR.json")
	if err != nil {
		t.Fatalf("Could not read JSON from filesystem: %s", err.Error())
	}

	var result Ticker
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Fatalf("Could not parse JSON: %s", err.Error())
	}

	if len(result) != 1 {
		t.Fatalf("Returned wrong length, got %d, expected 1", len(result))
	}

	coin := result.Index(0)

	usd, err := coin.Price("USD")
	if err != nil {
		t.Fatalf("Failed to get price in USD: %s", err.Error())
	}

	if usd != 18461.3 {
		t.Errorf("Price seems wrong, got %f, expected 18461.3", usd)
	}

	eur, err := coin.Price("eur") // lowercase for more testing
	if err != nil {
		t.Fatalf("Failed to get price in EUR: %s", err.Error())
	}

	if eur != 15716.0123835 {
		t.Errorf("Price seems wrong, got %f, expected 15716.0123835", eur)
	}
}

func TestCoinInfoAccessorsFail(t *testing.T) {
	var coin CoinInfo

	_, err := coin.Price("BOOM")
	if err == nil {
		t.Errorf("Price() did not fail for unknown currency symbol")
	}

	_, err = coin.Volume24H("BOOM")
	if err == nil {
		t.Errorf("Volume24H() did not fail for unknown currency symbol")
	}

	_, err = coin.MarketCap("BOOM")
	if err == nil {
		t.Errorf("Volume24H() did not fail for unknown currency symbol")
	}
}

func TestCoinInfoUnmarshalJSON(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ticker-1-EUR.json")
	if err != nil {
		t.Fatalf("Could not read JSON from filesystem: %s", err.Error())
	}

	var result Ticker
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Fatalf("Could not parse JSON: %s", err.Error())
	}

	if len(result) != 1 {
		t.Fatalf("Returned wrong length, got %d, expected 1", len(result))
	}

	coin := result.Index(0)
	if coin.Rank != 1 {
		t.Errorf("Read wrong rank from JSON")
	}
}

func TestCoinInfoUnmarshalJSONError(t *testing.T) {
	b := []byte(`{"ID":12}`)

	var coin CoinInfo
	err := json.Unmarshal(b, &coin)
	if err == nil {
		t.Fatalf("Did not catch broken JSON")
	}
}

func TestCoinInfoUnmarshalJSONError2(t *testing.T) {
	b := []byte(`{"IDD":12}`)

	var coin CoinInfo
	err := json.Unmarshal(b, &coin)
	if err == nil {
		t.Fatalf("Did not catch broken JSON")
	}
}

func TestCoinInfoUnmarshalJSONError3(t *testing.T) {
	b := []byte(`{"price_usd":"hejsa"}`)

	var coin CoinInfo
	err := json.Unmarshal(b, &coin)
	if err == nil {
		t.Fatalf("Did not catch broken JSON")
	}
}
