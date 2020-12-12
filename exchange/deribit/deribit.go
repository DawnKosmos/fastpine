package deribit

import (
	"github.com/frankrap/deribit-api"
)

type DERIBIT struct {
	d *deribit.Client
}

func New(key, secret string) *DERIBIT {
	cfg := &deribit.Configuration{
		Addr:      deribit.RealBaseURL,
		ApiKey:    key,
		SecretKey: secret,
	}
	client := DERIBIT{deribit.New(cfg)}

	return &client
}
