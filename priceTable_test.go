package coinmarketcap

import (
	"testing"
)

func TestPriceTableAdd(t *testing.T) {
	table := make(priceTable)
	err := table.add("price_usd", "12.2", "price_")
	if err != nil {
		t.Fatalf("add() failed on good input: %s", err.Error())
	}

	if table["USD"] != 12.2 {
		t.Errorf("USD is not 12.2, got %f", table["USD"])
	}
}

func TestPriceTableAddError(t *testing.T) {
	table := make(priceTable)
	err := table.add("price_usd", "not-float", "price_")
	if err == nil {
		t.Fatalf("add() failed to err")
	}
}

func TestPriceTableGet(t *testing.T) {
	table := make(priceTable)

	table["USD"] = 12.3

	value, err := table.get("USD")
	if err != nil {
		t.Fatalf("get() errored: %s", err.Error())
	}

	if value != 12.3 {
		t.Errorf("got %f, expected 12.3", value)
	}
}

func TestPriceTableGetError(t *testing.T) {
	table := make(priceTable)
	_, err := table.get("BOOM")
	if err == nil {
		t.Fatalf("get() failed to err")
	}
}
