package coinmarketcap

import (
	"encoding/json"
	"errors"
	"time"
)

// Point represents a value at a specific time.
type Point struct {
	Timestamp time.Time
	Value     float64
}

// UnmarshalJSON implements json.Unmarshaller. This is needed because the
// format returned from coinmarketcap is encoded as an array for each point.
func (p *Point) UnmarshalJSON(data []byte) error {
	proxy := make([]float64, 0, 2)

	err := json.Unmarshal(data, &proxy)
	if err != nil || len(proxy) != 2 {
		return errors.New("JSON is not in expected format")
	}

	p.Timestamp = time.Unix(int64(proxy[0]/1000.0), 0)
	p.Value = proxy[1]

	return nil
}
