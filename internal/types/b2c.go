package types

//B2Cpayment payment from buisness to customer
type B2Cpayment struct {
	InitiatorName      string `json:"InitiatorName"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	Amount             int    `json:"Amount"`
	PartyA             int    `json:"PartyA"`
	PartyB             int    `json:"PartyB"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
	Occassion          string `json:"Occassion"`
}
