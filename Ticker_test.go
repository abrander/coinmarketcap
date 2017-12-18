package coinmarketcap

import (
	"testing"
)

func TestTickerCoinByID(t *testing.T) {
	ticker := Ticker{
		&CoinInfo{ID: "bitcoin", Symbol: "BTC"},
		&CoinInfo{ID: "litecoin", Symbol: "LTC"},
	}

	if coin := ticker.CoinByID("bitcoin"); coin == nil || coin.Symbol != "BTC" {
		t.Errorf("CoinByID(\"bitcoin\") failed")
	}

	if ticker.CoinByID("boom") != nil {
		t.Errorf("CoinByID(\"boom\") failed to return nil")
	}
}

func TestTickerCoinBySymbol(t *testing.T) {
	ticker := Ticker{
		&CoinInfo{ID: "bitcoin", Symbol: "BTC"},
		&CoinInfo{ID: "litecoin", Symbol: "LTC"},
	}

	if coin := ticker.CoinBySymbol("LTC"); coin == nil || coin.ID != "litecoin" {
		t.Errorf("CoinBySymbol(\"LTC\") failed")
	}

	if ticker.CoinBySymbol("boom") != nil {
		t.Errorf("CoinBySymbol(\"boom\") failed to return nil")
	}
}

func TestTickerIndex(t *testing.T) {
	ticker := Ticker{
		&CoinInfo{ID: "bitcoin", Symbol: "BTC"},
		&CoinInfo{ID: "litecoin", Symbol: "LTC"},
	}

	if coin := ticker.Index(1); coin == nil || coin.ID != "litecoin" {
		t.Errorf("Index(1) failed")
	}

	if ticker.Index(2) != nil {
		t.Errorf("Index(2) failed to return nil")
	}
}
