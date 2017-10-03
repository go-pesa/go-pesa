package api

import (
	"fmt"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/http"
	"github.com/go-pesa/internal/types"
	"github.com/go-pesa/internal/utils"
)

//RequestBalance requests for balance of entity
func RequestBalance(party string, identifier int, remarks string) types.BalanceResponse {
	endpoint := balanceRequest
	pushObject := &types.BalanceRequest{
		CommandID:          "AccountBalance",
		PartyA:             party,
		IdentifierType:     identifier,
		Remarks:            remarks,
		Initiator:          globals.Initiator,
		SecurityCredential: globals.SecurityCredential,
		QueueTimeOutURL:    globals.QueueTimeOutURL,
		ResultURL:          globals.BalanceResultURL,
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

	var balanceResponse types.BalanceResponse

	utils.Unmarshal(res, &balanceResponse)

	return balanceResponse

}
