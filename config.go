package gopesa

import (
	"github.com/go-pesa/go-pesa/internal/api"
)

//Config struct, used for simplifying mobile and C wrappers
type Config struct {
	//ProductionURL is the base URl for all endpoints, if not defined defaults to sandbox
	ProductionURL string
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

//CreateConfig Creates an options configuration object
func (config *Config) CreateConfig() api.Options {
	var options api.Options

	if config.ProductionURL != "" {
		options.Add(ProductionURL(config.ProductionURL))
	}

	if config.ShortCode != 0 {
		options.Add(ShortCode(config.ShortCode))
	}

	if config.BusinessShortCode != 0 {
		options.Add(BusinessShortCode(config.BusinessShortCode))
	}

	if config.PassKey != "" {
		options.Add(PassKey(config.PassKey))
	}

	if config.TransactionCallback != "" {
		options.Add(TransactionCallback(config.TransactionCallback))
	}

	if config.MSISDN != "" {
		options.Add(MSISDN(config.MSISDN))
	}

	if config.Initiator != "" {
		options.Add(Initiator(config.Initiator))
	}

	return options
}

//MSISDN configures the MSISDN as an argument for New
func MSISDN(msisdn string) api.Option {
	return func(c *api.Client) {
		c.MSISDN = msisdn
	}
}

//ProductionURL configures the ProductionURL as an argument for New
func ProductionURL(url string) api.Option {
	return func(c *api.Client) {
		c.ProductionURL = url
	}
}

//ShortCode configures the ShortCode as an argument for New
func ShortCode(shortcode int) api.Option {
	return func(c *api.Client) {
		c.ShortCode = shortcode
	}
}

//BusinessShortCode configures the BusinessShortCode as an argument for New
func BusinessShortCode(businessShortCode int) api.Option {
	return func(c *api.Client) {
		c.BusinessShortCode = businessShortCode
	}
}

//PassKey configures the PassKey as an argument for New
func PassKey(passKey string) api.Option {
	return func(c *api.Client) {
		c.PassKey = passKey
	}
}

//TransactionCallback configures the TransactionCallback as an argument for New
func TransactionCallback(url string) api.Option {
	return func(c *api.Client) {
		c.TransactionCallback = url
	}
}

//Initiator configures the Initiator as an argument for New
func Initiator(code string) api.Option {
	return func(c *api.Client) {
		c.Initiator = code
	}
}
