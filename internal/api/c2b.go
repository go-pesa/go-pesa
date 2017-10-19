package api

// //RegisterC2BURL registers URLs for C2B
// func RegisterC2BURL(onfail string) types.GenericResponse {
// 	endpoint := registerC2BURL

// 	pushObject := &types.C2BURLRegisterObject{
// 		ShortCode:       globals.ShortCode,
// 		ResponseType:    onfail,
// 		ConfirmationURL: globals.ConfirmationURL,
// 		ValidationURL:   globals.ValidationURL,
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

// 	var balanceResponse types.GenericResponse

// 	utils.Unmarshal(res, &balanceResponse)

// 	return balanceResponse

// }

// //SimulateC2B simulates C2B
// func SimulateC2B(amount int, phoneNumber int, refNumber string, command string) types.GenericResponse {
// 	endpoint := simulateC2B

// 	pushObject := &types.C2BSimulationObject{
// 		ShortCode:     globals.ShortCode,
// 		Amount:        amount,
// 		MSISDN:        phoneNumber,
// 		BillRefNumber: refNumber,
// 		CommandID:     command,
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

// 	var simulationResponse types.GenericResponse

// 	utils.Unmarshal(res, &simulationResponse)

// 	return simulationResponse

// }
