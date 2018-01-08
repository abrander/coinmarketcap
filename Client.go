package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is a client for talking to the coinmarketcap API.
type Client struct {
	baseURL          string
	graphBaseURL     string
	requestPerMinute int
	ratelimit        chan bool
	quit             chan bool
	client           *http.Client
}

// BaseURL will define a new base URL. You would probably never use this.
func BaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// GraphBaseURL sets the base URL for graph data. You will probablye never
// need this.
func GraphBaseURL(graphBaseURL string) func(*Client) {
	return func(c *Client) {
		c.graphBaseURL = graphBaseURL
	}
}

// RateLimit will set at which limit requests should be throttled. Default is
// 10 per minute as kindly requested by CoinMarketCap.
func RateLimit(requestPerMinute int) func(*Client) {
	return func(c *Client) {
		c.requestPerMinute = requestPerMinute
	}
}

// HTTPClient will change the HTTP client to use. Default is
// http.DefaultClient. This can be used for example if you would like to use
// the AppEngine http client.
func HTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.client = client
	}
}

// NewClient will return a new client. For now this will never return an error,
// but you should check it anyway. Maybe some time in the future we will
// return an error.
func NewClient(options ...func(*Client)) (*Client, error) {
	client := &Client{
		baseURL:          "https://api.coinmarketcap.com/v1",
		graphBaseURL:     "https://graphs.coinmarketcap.com/",
		quit:             make(chan bool),
		requestPerMinute: 10,
		client:           http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	client.ratelimit = make(chan bool, client.requestPerMinute)

	go client.throttler()

	return client, nil
}

func (c *Client) throttler() {
	// Allow the client to use all requests at once.
	for i := 0; i < c.requestPerMinute; i++ {
		c.ratelimit <- true
	}

	ticker := time.NewTicker(time.Minute / time.Duration(c.requestPerMinute))

	for {
		select {
		case <-ticker.C:
			select {
			case c.ratelimit <- true:
			default:
			}
		case <-c.quit:
			ticker.Stop()
			return
		}
	}
}

// Close will close the API client. For now this is a no-op. For future
// compatibility you should always close the client.
func (c *Client) Close() error {
	c.quit <- true

	return nil
}

// Ticker calls a ticker API endpoint and returns a Ticker which you can
// range through. Options are defined with the functions Start(), Limit(),
// Convert() and Currency().
func (c *Client) Ticker(options ...func(*query)) (Ticker, error) {
	q := newQuery(options)

	URL := c.baseURL + "/ticker/" + q.tickerQuery()

	resp, err := c.get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Ticker

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

	resp, err := c.get(URL)
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

// Graph will retrieve graph data. This is NOT part of the official
// CoinMarketCap API. Currency could be "bitcoin".
func (c *Client) Graph(currency string, from time.Time, to time.Time) (*GraphData, error) {
	URL := fmt.Sprintf("%s/currencies/%s/%d/%d", c.graphBaseURL, currency, from.Unix()*1000, to.Unix()*1000)

	resp, err := c.get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GraphData

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) get(URL string) (*http.Response, error) {
	<-c.ratelimit

	return c.client.Get(URL)
}
