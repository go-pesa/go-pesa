package types

//C2BURLRegisterObject used to create validation and confirmation URLS
type C2BURLRegisterObject struct {
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

//C2BSimulationObject used for simulating transactions
type C2BSimulationObject struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        int    `json:"Amount"`
	MSISDN        int    `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}
