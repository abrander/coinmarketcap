package coinmarketcap

import (
	"fmt"
	"net/url"
)

// query is used for defining what to lookup
type query struct {
	start    *int
	limit    *int
	convert  string
	currency string
}

func newQuery(options []func(*query)) *query {
	q := &query{}

	for _, option := range options {
		option(q)
	}

	return q
}

// Start will start the query at rank start+1.
func Start(start int) func(*query) {
	return func(q *query) {
		q.start = &start
	}
}

// Limit makes sure the query only returns limit entries.
func Limit(limit int) func(*query) {
	return func(q *query) {
		q.limit = &limit
	}
}

// Convert will make sure the result is also available in currency. An example
// could be "DKK". Prices will always be returned in USD and BTC.
func Convert(currency string) func(*query) {
	return func(q *query) {
		q.convert = currency
	}
}

// Currency will only look up currency. Please note that this is NOT the
// symbol. An example could be "bitcoin" or "litecoin".
func Currency(currency string) func(*query) {
	return func(q *query) {
		q.currency = currency
	}
}

// Return URL parameters representing this query.
func (q *query) tickerQuery() string {
	if q.start == nil && q.limit == nil && q.convert == "" {
		return q.currency
	}

	values := &url.Values{}

	if q.start != nil {
		values.Add("start", fmt.Sprintf("%d", *q.start))
	}

	if q.limit != nil {
		values.Add("limit", fmt.Sprintf("%d", *q.limit))
	}

	if q.convert != "" {
		values.Add("convert", q.convert)
	}

	return q.currency + "?" + values.Encode()
}
