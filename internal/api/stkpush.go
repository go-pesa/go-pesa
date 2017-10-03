package api

import (
	"fmt"
	"time"

	"github.com/go-pesa/internal/globals"

	"github.com/go-pesa/internal/http"
	"github.com/go-pesa/internal/types"
	"github.com/go-pesa/internal/utils"
)

//StkPushProcess creates an stk push
func StkPushProcess(amount int, customerPhone int, ref string, desc string) types.StkResponse {
	timestamp := time.Now().Format("20060102150405")
	endpoint := stkpushProcessRequest
	password := utils.EncodePassword(globals.ShortCode, globals.PassKey, timestamp)
	pushObject := &types.StkPush{
		BusinessShortCode: globals.ShortCode,
		Timestamp:         timestamp,
		Password:          password,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            globals.MSISDN,
		PartyB:            globals.ShortCode,
		PhoneNumber:       customerPhone,
		CallBackURL:       globals.Callback,
		AccountReference:  ref,
		TransactionDesc:   desc,
	}
	var headers = []types.Header{
		types.Header{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken()),
		},
		types.Header{
			Key:   "content-type",
			Value: "application/json",
		},
		types.Header{
			Key:   "cache-control",
			Value: "no-cache",
		},
	}

	res := http.Post(endpoint, headers, pushObject)

	var stkresponse types.StkResponse

	utils.Unmarshal(res, &stkresponse)

	return stkresponse
}

//StkPushQuery queries an stk push
func StkPushQuery(CheckoutRequestID string) types.StkResponse {
	timestamp := time.Now().Format("20060102150405")
	endpoint := stkpushQuery
	password := utils.EncodePassword(globals.ShortCode, globals.PassKey, timestamp)
	pushObject := &types.StkPushQuery{
		BusinessShortCode: globals.ShortCode,
		Password:          password,
		Timestamp:         timestamp,
		CheckoutRequestID: CheckoutRequestID,
	}
	var headers = []types.Header{
		types.Header{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken()),
		},
		types.Header{
			Key:   "content-type",
			Value: "application/json",
		},
		types.Header{
			Key:   "cache-control",
			Value: "no-cache",
		},
	}

	res := http.Post(endpoint, headers, pushObject)

	var stkresponse types.StkResponse

	utils.Unmarshal(res, &stkresponse)

	return stkresponse
}
