package types

//B2BPaymentRequest is the request object for B2B payments
type B2BPaymentRequest struct {
	CommandID              string `json:"CommandID"`
	Amount                 int    `json:"Amount"`
	PartyA                 int    `json:"PartyA"`
	SenderIdentifierType   int    `json:"SenderIdentifierType"`
	PartyB                 int    `json:"PartyB"`
	RecieverIdentifierType int    `json:"RecieverIdentifierType"`
	Remarks                string `json:"Remarks"`
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	AccountReference       string `json:"AccountReference"`
}
