package gopesa

import (
	"github.com/bluele/gcache"
	"github.com/go-pesa/go-pesa/internal/api"
)

//New creates a new client object
func New(key, secret string, opts ...api.Option) *api.Client {

	c := &api.Client{
		MSISDN:              254708374149,
		Secret:              secret,
		Key:                 key,
		ProductionURL:       "https://sandbox.safaricom.co.ke/",
		PassKey:             "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
		TransactionCallback: "https://pay-pesa.firebaseio.com/pay.json",
		ShortCode:           174379,
		Cache:               gcache.New(20).LRU().Build(),
		Initiator:           "apiop71",
		BusinessShortCode:   602948,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
