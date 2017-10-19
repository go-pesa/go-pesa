package api

import (
	"log"
)

func check(reqs string, client *Client) bool {

	var res bool
	switch reqs {

	case "stkpush":
		res = checkStkPushRequirements(client)
	default:
		res = false
	}

	return res

}

func checkStkPushRequirements(client *Client) bool {

	if client.Key == "" ||
		client.PassKey == "" ||
		client.TransactionCallback == "" ||
		client.ShortCode == 0 ||
		client.MSISDN == 0 {
		log.Println(`Please make sure you have instantiated the client with the following properties\n
						 1. Key (Consumer Key)\n
						 2. TransactionCallback (Callback URL for STK push)\n
						 3. ShortCode (Lipa Na M-Pesa ShortCode)\n
						 4. MSISDN (MSISDN - phone number provided for initiating transactions)`)
		return false
	}

	return true
}
