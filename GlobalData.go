package coinmarketcap

import (
	"encoding/json"
	"strings"
	"time"
)

// GlobalData will describe all cryptocurrencies known on coinmarketcap.
type GlobalData struct {
	BitcoinPercentage float64   `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies  int       `json:"active_currencies"`
	ActiveAssets      int       `json:"active_assets"`
	ActiveMarkets     int       `json:"active_markets"`
	LastUpdated       time.Time `json:"last_updated"`

	volume24h priceTable
	cap       priceTable
}

// Volume24H returns the traded volume in the last 24 hours. A currency must
// be provided.
func (g *GlobalData) Volume24H(currency string) (float64, error) {
	return g.volume24h.get(currency)
}

// MarketCap returns the market cap. A currency must be provided.
func (g *GlobalData) MarketCap(currency string) (float64, error) {
	return g.cap.get(currency)
}

// UnmarshalJSON implements json.Unmarshaler. We have to implement this because
// the json from the coinmarketcap API is not easily parsable using Go's JSON
// unmarshaller..
func (g *GlobalData) UnmarshalJSON(b []byte) error {
	var proxy struct {
		BitcoinPercentage float64 `json:"bitcoin_percentage_of_market_cap"`
		ActiveCurrencies  int     `json:"active_currencies"`
		ActiveAssets      int     `json:"active_assets"`
		ActiveMarkets     int     `json:"active_markets"`
		LastUpdated       int64   `json:"last_updated"`
	}

	err := json.Unmarshal(b, &proxy)
	if err != nil {
		return err
	}

	g.BitcoinPercentage = proxy.BitcoinPercentage
	g.ActiveCurrencies = proxy.ActiveCurrencies
	g.ActiveAssets = proxy.ActiveAssets
	g.ActiveMarkets = proxy.ActiveMarkets
	g.LastUpdated = time.Unix(proxy.LastUpdated, 0)

	g.volume24h = make(priceTable)
	g.cap = make(priceTable)

	var kv map[string]json.RawMessage

	// This should never fail.
	json.Unmarshal(b, &kv)

	for key, value := range kv {
		switch {
		case strings.HasPrefix(key, "total_24h_volume_"):
			err = g.volume24h.add(key, string(value), "total_24h_volume_")

		case strings.HasPrefix(key, "total_market_cap_"):
			err = g.cap.add(key, string(value), "total_market_cap_")
		}

		if err != nil {
			return err
		}
	}

	return err
}
