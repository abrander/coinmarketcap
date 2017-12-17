# coinmarketcap

This package provides access to the public [CoinMarketCap](https://coinmarketcap.com/) [API](https://coinmarketcap.com/api/).

[![GoDoc][1]][2]

[1]: https://godoc.org/github.com/abrander/coinmarketcap?status.svg
[2]: https://godoc.org/github.com/abrander/coinmarketcap

## Overview

This package uses the [functional options pattern](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
to ensure that we can upgrade this package without breaking compatibility if the public API ever changes. 

## Examples

### Printing the global market cap in danish kroner

```golang
import (
	"fmt"
	"time"

	"github.com/abrander/coinmarketcap"
)

func main() {
	client, _ := coinmarketcap.NewClient()

	globaldata, _ := client.GlobalData(
		coinmarketcap.Convert("DKK"),
	)

	cap, _ := globaldata.MarketCap("DKK")

	fmt.Printf("Global market cap in DKK: %.0f (Updated %s ago)\n",
		cap,
		time.Since(globaldata.LastUpdated),
	)
}
```

### Get the top five changes

```go
import (
	"fmt"

	"github.com/abrander/coinmarketcap"
)

func main() {
	client, _ := coinmarketcap.NewClient()

	coins, _ := client.Ticker(
		coinmarketcap.Limit(5),
	)

	for _, coin := range coins {
		fmt.Printf("%04s %s 1H: %f 24H:%f  7D:%f (%s)\n",
			coin.Symbol,
			coin.ID,
			coin.PercentChange1H,
			coin.PercentChange24H,
			coin.PercentChange7D,
			coin.LastUpdated,
		)
	}
}
```
