package bybit

import (
	"net/http"

	"github.com/frankrap/bybit-api/rest"
)

type BYBIT struct {
	b *rest.ByBit
}

const BASEURL string = "https://api.bybit.com/"
const TESTURL string = "https://api-testnet.bybit.com/"

func New(test bool, Client *http.Client, public string, secret string) *BYBIT {
	baseUrl := BASEURL
	if test {
		baseUrl = TESTURL
	}
	var b BYBIT
	b.b = rest.New(Client, baseUrl, "", "", false)
	return &b
}
