package api

// //MakeB2BPayment does a B2B payment
// func MakeB2BPayment(amount int, receiver int, receiverType int, sender int, senderType int, command string, reference string, remarks string) types.GenericResponse {
// 	endpoint := processB2B

// 	pushObject := &types.B2BPaymentRequest{
// 		CommandID:              command,
// 		Amount:                 amount,
// 		PartyA:                 sender,
// 		SenderIdentifierType:   senderType,
// 		PartyB:                 receiver,
// 		RecieverIdentifierType: receiverType,
// 		Remarks:                remarks,
// 		Initiator:              globals.Initiator,
// 		SecurityCredential:     globals.SecurityCredential,
// 		QueueTimeOutURL:        globals.QueueTimeOutURL,
// 		ResultURL:              globals.BalanceResultURL,
// 		AccountReference:       reference,
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

// 	var b2bResponse types.GenericResponse

// 	utils.Unmarshal(res, &b2bResponse)

// 	return b2bResponse

// }
