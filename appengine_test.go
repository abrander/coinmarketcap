// +build go1.8

package coinmarketcap

import (
	"context"
	"testing"

	"google.golang.org/appengine/urlfetch"
)

func TestAppEngineCompatibility(t *testing.T) {
	// If this compiles we should be good to go ;)
	NewClient(HTTPClient(urlfetch.Client(context.TODO())))
}
