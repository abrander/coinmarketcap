package coinmarketcap

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrCurrencyNotFound will be returned if the requested currency is not available.
	ErrCurrencyNotFound = errors.New("currency not found")
)

// priceTable can hold values in different currencies. It's useful when the
// API returns values in different currencies.
type priceTable map[string]float64

// add a new currency and value to t.
func (t priceTable) add(key string, value string, prefix string) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	symbol := strings.TrimPrefix(key, prefix)

	symbol = strings.ToUpper(symbol)

	t[symbol] = f

	return nil
}

// get the value for a specific symbol.
func (t priceTable) get(symbol string) (float64, error) {
	symbol = strings.ToUpper(symbol)

	value, found := t[symbol]
	if !found {
		return 0.0, ErrCurrencyNotFound
	}

	return value, nil
}
