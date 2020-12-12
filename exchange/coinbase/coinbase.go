package coinbase

import (
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

type COINBASE struct {
	Client *coinbasepro.Client

	key    string
	secret string
}

const URL = "https://api.pro.coinbase.com"

func New(key, secret string) *COINBASE {
	return &COINBASE{coinbasepro.NewClient(), key, secret}
}
