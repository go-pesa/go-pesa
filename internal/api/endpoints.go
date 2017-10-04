package api

var (
	oauthGenerate         = "oauth/v1/generate"
	stkpushQuery          = "mpesa/stkpushquery/v1/query"
	stkpushProcessRequest = "mpesa/stkpush/v1/processrequest"
	balanceRequest        = "mpesa/accountbalance/v1/query"
	registerC2BURL        = "mpesa/c2b/v1/registerurl"
	simulateC2B           = "mpesa/c2b/v1/simulate"
	processB2B            = "/mpesa/b2b/v1/paymentrequest"
	processC2B            = "/mpesa/b2c/v1/paymentrequest"
)
