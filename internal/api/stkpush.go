package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-pesa/go-pesa/internal/http"
	"github.com/go-pesa/go-pesa/internal/types"
	"github.com/go-pesa/go-pesa/internal/utils"
)

//StkPush creates an stk push(charges a customer by creating a SIM tool kit pop-up)
func (client *Client) StkPush(amount, customerPhone int, ref, desc string) types.StkResponse {
	var stkresponse types.StkResponse

	if check("stkpush", client) {
		return stkresponse
	}

	timestamp := time.Now().Format("20060102150405")
	endpoint := stkpushProcessRequest
	businessShortCode := strconv.Itoa(client.ShortCode)
	password := utils.EncodePassword(businessShortCode, client.PassKey, timestamp)

	pushObject := &types.StkPush{
		BusinessShortCode: businessShortCode,
		Timestamp:         timestamp,
		Password:          password,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            fmt.Sprintf("%v", client.MSISDN),
		PartyB:            businessShortCode,
		PhoneNumber:       customerPhone,
		CallBackURL:       client.TransactionCallback,
		AccountReference:  ref,
		TransactionDesc:   desc,
	}
	var headers = []types.Header{
		types.Header{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken(client.Key, client.Secret, client.ProductionURL, client.Cache)),
		},
	}

	res := http.Post(endpoint, client.ProductionURL, headers, pushObject)

	utils.Unmarshal(res, &stkresponse)

	return stkresponse
}

//StkPushQuery queries an stk push
func (client *Client) StkPushQuery(CheckoutRequestID string) types.StkResponse {

	var stkresponse types.StkResponse
	if check("stkpush", client) {
		return stkresponse
	}
	timestamp := time.Now().Format("20060102150405")
	endpoint := stkpushQuery
	password := utils.EncodePassword(fmt.Sprintf("%v", client.ShortCode), client.PassKey, timestamp)
	pushObject := &types.StkPushQuery{
		BusinessShortCode: fmt.Sprintf("%v", client.ShortCode),
		Password:          password,
		Timestamp:         timestamp,
		CheckoutRequestID: CheckoutRequestID,
	}
	var headers = []types.Header{
		types.Header{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken(client.Key, client.Secret, client.ProductionURL, client.Cache)),
		},
	}

	res := http.Post(endpoint, client.ProductionURL, headers, pushObject)

	utils.Unmarshal(res, &stkresponse)

	return stkresponse
}
