package gopesa

import "github.com/go-pesa/go-pesa/internal/api"

//MSISDN configures the MSISDN as an argument for New
func MSISDN(msisdn int64) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.MSISDN
		c.MSISDN = msisdn
		return MSISDN(previous)
	}
}

//ProductionURL configures the ProductionURL as an argument for New
func ProductionURL(url string) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.ProductionURL
		c.ProductionURL = url
		return ProductionURL(previous)
	}
}

//ShortCode configures the ShortCode as an argument for New
func ShortCode(shortcode int) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.ShortCode
		c.ShortCode = shortcode
		return ShortCode(previous)
	}
}

//BusinessShortCode configures the BusinessShortCode as an argument for New
func BusinessShortCode(businessShortCode int) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.BusinessShortCode
		c.BusinessShortCode = businessShortCode
		return BusinessShortCode(previous)
	}
}

//PassKey configures the PassKey as an argument for New
func PassKey(passKey string) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.PassKey
		c.PassKey = passKey
		return PassKey(previous)
	}
}

//TransactionCallback configures the TransactionCallback as an argument for New
func TransactionCallback(url string) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.TransactionCallback
		c.TransactionCallback = url
		return TransactionCallback(previous)
	}
}

//Initiator configures the Initiator as an argument for New
func Initiator(code string) api.Option {
	return func(c *api.Client) api.Option {
		previous := c.Initiator
		c.Initiator = code
		return Initiator(previous)
	}
}
