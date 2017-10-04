package api

import (
	"fmt"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/http"
	"github.com/go-pesa/internal/types"
	"github.com/go-pesa/internal/utils"
)

//MakeB2CPayment does a B2C payment
func MakeB2CPayment(amount int, phoneNumber int, sender int, command string, remarks string, occasion string) types.GenericResponse {
	endpoint := processB2B

	pushObject := &types.B2Cpayment{
		CommandID:          command,
		Amount:             amount,
		PartyA:             sender,
		PartyB:             phoneNumber,
		Remarks:            remarks,
		InitiatorName:      globals.Initiator,
		SecurityCredential: globals.SecurityCredential,
		QueueTimeOutURL:    globals.QueueTimeOutURL,
		ResultURL:          globals.BalanceResultURL,
		Occassion:          occasion,
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

	var b2bResponse types.GenericResponse

	utils.Unmarshal(res, &b2bResponse)

	return b2bResponse

}
