package api

import (
	"fmt"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/http"
	"github.com/go-pesa/internal/types"
	"github.com/go-pesa/internal/utils"
)

//CheckStatus does a B2C payment
func CheckStatus(transactionID string, sender int, senderType int, remarks string, occasion string) types.GenericResponse {
	endpoint := transactionStatus

	pushObject := &types.StatusRequest{
		CommandID:          "TransactionStatusQuery",
		PartyA:             sender,
		IdentifierType:     senderType,
		Remarks:            remarks,
		Initiator:          globals.Initiator,
		SecurityCredential: globals.SecurityCredential,
		QueueTimeOutURL:    globals.QueueTimeOutURL,
		ResultURL:          globals.BalanceResultURL,
		TransactionID:      transactionID,
		Occasion:           occasion,
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

	var statusRes types.GenericResponse

	utils.Unmarshal(res, &statusRes)

	return statusRes

}
