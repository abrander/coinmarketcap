package coinmarketcap

// GraphData represents multiple series of Points as returned by the graph API.
type GraphData struct {
	MarkepCap []Point `json:"market_cap_by_available_supply"`
	PriceBTC  []Point `json:"price_btc"`
	PriceUSD  []Point `json:"price_usd"`
	VolumeUSD []Point `json:"volume_usd"`
}
