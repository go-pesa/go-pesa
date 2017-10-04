package types

//BalanceRequest is the request object
type BalanceRequest struct {
	CommandID          string `json:"CommandID"`
	PartyA             string `json:"PartyA"`
	IdentifierType     int    `json:"IdentifierType"`
	Remarks            string `json:"Remarks"`
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
}
