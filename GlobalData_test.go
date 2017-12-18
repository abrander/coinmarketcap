package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGlobalDataAccessors(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/global.json")
	if err != nil {
		t.Fatalf("Could not read JSON from filesystem: %s", err.Error())
	}

	var data GlobalData
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("Could not parse JSON: %s", err.Error())
	}

	usd, err := data.MarketCap("USD")
	if err != nil {
		t.Fatalf("Failed to get market cap in USD: %s", err.Error())
	}

	if usd != 603106469791.0 {
		t.Errorf("Price seems wrong, got %f, expected 603106469791.0", usd)
	}

	usd, err = data.Volume24H("usd") // lowercase for more testing
	if err != nil {
		t.Fatalf("Failed to get volume in USD: %s", err.Error())
	}

	if usd != 30512869361.0 {
		t.Errorf("Price seems wrong, got %f, expected 30512869361.0", usd)
	}
}

func TestGlobalDataUnmarshalJSONError(t *testing.T) {
	b := []byte(`{"active_assets": "boom"}`)

	var coin GlobalData
	err := json.Unmarshal(b, &coin)
	if err == nil {
		t.Fatalf("Did not catch broken JSON")
	}
}

func TestGlobalDataUnmarshalJSONError2(t *testing.T) {
	b := []byte(`{"total_24h_volume_usd":"boom!"}`)

	var coin GlobalData
	err := json.Unmarshal(b, &coin)
	if err == nil {
		t.Fatalf("Did not catch broken JSON")
	}
}
