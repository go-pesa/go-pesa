package api

// //ReverseTransaction does a B2C payment
// func ReverseTransaction(transactionID string, amount int, receiver int, receiverType int, remarks string, occasion string) types.GenericResponse {
// 	endpoint := reverseTransaction

// 	pushObject := &types.ReverseTranactionRequest{
// 		CommandID:              "TransactionReversal",
// 		ReceiverParty:          receiver,
// 		RecieverIdentifierType: receiverType,
// 		Amount:                 amount,
// 		Remarks:                remarks,
// 		Initiator:              globals.Initiator,
// 		SecurityCredential:     globals.SecurityCredential,
// 		QueueTimeOutURL:        globals.QueueTimeOutURL,
// 		ResultURL:              globals.BalanceResultURL,
// 		TransactionID:          transactionID,
// 		Occasion:               occasion,
// 	}
// 	var headers = []types.Header{
// 		types.Header{
// 			Key:   "authorization",
// 			Value: fmt.Sprintf("Bearer %s", getToken()),
// 		},
// 		types.Header{
// 			Key:   "content-type",
// 			Value: "application/json",
// 		},
// 		types.Header{
// 			Key:   "cache-control",
// 			Value: "no-cache",
// 		},
// 	}

// 	res := http.Post(endpoint, headers, pushObject)

// 	var statusRes types.GenericResponse

// 	utils.Unmarshal(res, &statusRes)

// 	return statusRes

// }
