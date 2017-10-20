package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-pesa/go-pesa/internal/http"
	"github.com/go-pesa/go-pesa/internal/types"
	"github.com/go-pesa/go-pesa/internal/utils"
)

//CheckIdentity checks the user's identity
func (client *Client) CheckIdentity(customerPhone int) types.StkResponse {
	var stkresponse types.StkResponse

	// if check("stkpush", client) {
	// 	return stkresponse
	// }

	timestamp := time.Now().Format("20060102150405")
	endpoint := identityCheck
	businessShortCode := strconv.Itoa(client.BusinessShortCode)
	password := utils.EncodePassword(businessShortCode, client.PassKey, timestamp)

	pushObject := &types.IdentityRequest{

		Initiator:         client.Initiator,
		BusinessShortCode: businessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CheckIdentity",
		PhoneNumber:       customerPhone,
		CallBackURL:       client.TransactionCallback,
		TransactionDesc:   "",
	}
	var headers = []types.Header{
		types.Header{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken(client.Key, client.Secret, client.ProductionURL, client.cache)),
		},
	}

	res := http.Post(endpoint, client.ProductionURL, headers, pushObject)

	utils.Unmarshal(res, &stkresponse)

	return stkresponse
}
