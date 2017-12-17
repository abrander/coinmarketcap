package coinmarketcap

import (
	"strings"
)

// TickerResult will be returned from Ticker().
type TickerResult []*CoinInfo

// CoinByID will search for a coin identified by ID. An example could be
// "bitcoin".
func (r TickerResult) CoinByID(ID string) *CoinInfo {
	for _, coin := range r {
		if coin.ID == ID {
			return coin
		}
	}

	return nil
}

// CoinBySymbol will search for a coin by its symbol. For example "BTC".
func (r TickerResult) CoinBySymbol(symbol string) *CoinInfo {
	symbol = strings.ToUpper(symbol)

	for _, coin := range r {
		if coin.Symbol == symbol {
			return coin
		}
	}

	return nil
}

// Index returns the coin at position n or nil if not found.
func (r TickerResult) Index(n int) *CoinInfo {
	if n >= len(r) {
		return nil
	}

	return r[n]
}
