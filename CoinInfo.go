package coinmarketcap

import (
	"encoding/json"
	"strings"
	"time"
)

// CoinInfo describes a coin.
type CoinInfo struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Rank             int       `json:"rank"`
	AvailableSupply  float64   `json:"available_supply"`
	TotalSupply      float64   `json:"total_supply"`
	MaxSupply        float64   `json:"max_supply"`
	PercentChange1H  float64   `json:"percent_change_1h"`
	PercentChange24H float64   `json:"percent_change_24h"`
	PercentChange7D  float64   `json:"percent_change_7d"`
	LastUpdated      time.Time `json:"last_updated"`

	price     priceTable
	volume24h priceTable
	cap       priceTable
}

// Price returns the price of the coin. currency must be used to specify in
// which currency the price is returned. If you use anything but "USD" and
// "BTC", you must use the Convert() function for Ticker().
func (i *CoinInfo) Price(currency string) (float64, error) {
	return i.price.get(currency)
}

// Volume24H returns the traded volume in the last 24 hours. A currency must
// be provided.
func (i *CoinInfo) Volume24H(currency string) (float64, error) {
	return i.volume24h.get(currency)
}

// MarketCap returns the market cap. A currency must be provided. This should
// match Price(currency) * AvailableSupply.
func (i *CoinInfo) MarketCap(currency string) (float64, error) {
	return i.cap.get(currency)
}

// UnmarshalJSON implements json.Unmarshaler. We have to implement this because
// the json from the coinmarketcap API is not easily parsable.
func (i *CoinInfo) UnmarshalJSON(b []byte) error {
	var proxy struct {
		ID               string  `json:"id"`
		Name             string  `json:"name"`
		Symbol           string  `json:"symbol"`
		Rank             int     `json:"rank,string"`
		AvailableSupply  float64 `json:"available_supply,string"`
		TotalSupply      float64 `json:"total_supply,string"`
		MaxSupply        float64 `json:"max_supply,string"`
		PercentChange1H  float64 `json:"percent_change_1h,string"`
		PercentChange24H float64 `json:"percent_change_24h,string"`
		PercentChange7D  float64 `json:"percent_change_7d,string"`
		LastUpdated      int64   `json:"last_updated,string"`
	}

	err := json.Unmarshal(b, &proxy)
	if err != nil {
		return err
	}

	i.ID = strings.ToLower(proxy.ID)
	i.Name = proxy.Name
	i.Symbol = strings.ToUpper(proxy.Symbol)
	i.Rank = proxy.Rank
	i.AvailableSupply = proxy.AvailableSupply
	i.TotalSupply = proxy.TotalSupply
	i.MaxSupply = proxy.MaxSupply
	i.PercentChange1H = proxy.PercentChange1H
	i.PercentChange24H = proxy.PercentChange24H
	i.PercentChange7D = proxy.PercentChange7D
	i.LastUpdated = time.Unix(proxy.LastUpdated, 0)

	i.price = make(priceTable)
	i.volume24h = make(priceTable)
	i.cap = make(priceTable)

	var kv map[string]string

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for key, value := range kv {
		switch {
		case strings.HasPrefix(key, "price_"):
			err = i.price.add(key, value, "price_")

		case strings.HasPrefix(key, "24h_volume_"):
			err = i.volume24h.add(key, value, "24h_volume_")

		case strings.HasPrefix(key, "market_cap_"):
			err = i.cap.add(key, value, "market_cap_")
		}

		if err != nil {
			return err
		}
	}

	return err
}
