package coinmarketcap

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPointFromJSON(t *testing.T) {
	p := Point{}
	b := []byte("[1506430755000, 65544642358]")

	err := json.Unmarshal(b, &p)
	if err != nil {
		t.Fatalf("Unmarshal failed: %s", err.Error())
	}

	if p.Value != 65544642358.0 {
		t.Errorf("Wrong value decoded, expected 65544642358.0, got %f", p.Value)
	}

	if p.Timestamp != time.Unix(1506430755, 0) {
		t.Errorf("Wrong timestamp decoded, expected %s, got %s", time.Unix(1506430755, 0), p.Timestamp)
	}
}

func TestPointFromBrokenJSON(t *testing.T) {
	cases := []string{
		"--",
		"[]",
		"[0.0,0,0]",
		"[0.0]",
		"[,]",
		"",
		"{",
		"[[[",
		"[2.0, 1.0, 0.0]",
	}

	for i, c := range cases {
		p := Point{}
		err := p.UnmarshalJSON([]byte(c))
		if err == nil {
			t.Errorf("%d Unmarshaller failed to detect broken JSON: %s", i, c)
		}
	}
}
