package gopesa

import (
	"github.com/go-pesa/go-pesa/internal/api"
)

//Client export for gomobile

//New creates a new client object
func New(key, secret string, opts ...api.Option) *api.Client {

	c := &api.Client{
		MSISDN:              "254708374149",
		Secret:              secret,
		Key:                 key,
		ProductionURL:       "https://sandbox.safaricom.co.ke/",
		PassKey:             "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
		TransactionCallback: "https://sandbox.safaricom.co.ke/",
		ShortCode:           174379,
		Initiator:           "apiop71",
		BusinessShortCode:   602948,
	}

	c.BuildCache()

	for _, opt := range opts {
		opt(c)
	}

	return c
}

//CastClient type asserts interface as Client object
func EmptyClient() *api.Client {
	return &api.Client{}
}

//Options creates an array of configuration oprions
func Options() *api.Options {
	var configs *api.Options
	return configs
}
