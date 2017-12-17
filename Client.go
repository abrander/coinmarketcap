package coinmarketcap

import (
	"encoding/json"
	"net/http"
)

// Client is a client for talking to the coinmarketcap API.
type Client struct {
	baseURL string
}

// BaseURL will define a new base URL. You would probably never use this.
func BaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient will return a new client. For now this will never return an error,
// but you should check it anyway. Maybe some time in the future we will
// return an error.
func NewClient(options ...func(*Client)) (*Client, error) {
	client := &Client{
		baseURL: "https://api.coinmarketcap.com/v1",
	}

	for _, option := range options {
		option(client)
	}

	return client, nil
}

// Ticker calls a ticker API endpoint and returns a TickerResult which you
// can range through. Options are defined with the functions Start(), Limit(),
// Convert() and Currency().
func (c *Client) Ticker(options ...func(*query)) (TickerResult, error) {
	q := newQuery(options)

	URL := c.baseURL + "/ticker/" + q.tickerQuery()

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TickerResult

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GlobalData returns global coin statistics from coinmarketcap.
func (c *Client) GlobalData(options ...func(*query)) (*GlobalData, error) {
	q := newQuery(options)

	URL := c.baseURL + "/global/" + q.tickerQuery()

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GlobalData

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
