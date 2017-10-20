package api

import (
	"github.com/bluele/gcache"
)

//Client is the main interface for the API
type Client struct {
	//ProductionURL is the base URl for all endpoints, if not defined defaults to sandbox
	ProductionURL string
	//Secret is the mpesa consumer secret, please set it with envirnment variables as MPESA_CONSUMER_SECRET
	Secret string
	//Key is the mpesa consumer key, please set it with envirnment variables as MPESA_CONSUMER_KEY
	Key string
	//shortCode for LIPA na M-PESA
	ShortCode int
	//BusinessShortcode is porvided by Safaricom(May be the same as LIPA NA M-PESA shortcode)
	BusinessShortCode int
	//PassKey is provided by M-pesa
	PassKey string
	//TransactionCallback is the Callback url for mpesa transactions
	TransactionCallback string
	//MSISDN is the MSISDN provided by Safaricom
	MSISDN string
	//Cache to hold temporary info like auth tokens
	cache gcache.Cache
	//Initiator as defined by the docmentation
	Initiator string
	// //InitiatorShortCode is the shortcode for initiator transactions
	// InitiatorShortCode string
	// //MpesaSecret
	// mpesaSecret string
	// //SecurityCredential encrypted MpesaSecret
	// SecurityCredential string
	// //QueueTimeOutURL M-pesa queue timeout URL
	// QueueTimeOutURL string
	// //BalanceResultURL M-pesa balance result URL
	// BalanceResultURL string
	// //ValidationURL c2B validation URL
	// ValidationURL string
	// //ConfirmationURL c2B confirmation URL
	// ConfirmationURL string
	// //Verbosity
	// x int

}

//Option is used to configure the API
type Option func(c *Client)

//Options is an array of client options
type Options []Option

//Add appends an option to the Options
func (opts *Options) Add(opt Option) Options {
	return append(*opts, opt)
}

//BuildCache creates the cache object
func (client *Client) BuildCache() {
	client.cache = gcache.New(20).LRU().Build()
}
