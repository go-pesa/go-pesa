package types

//StatusRequest used for checking transaction status
type StatusRequest struct {
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"CommandID"`
	TransactionID            string `json:"TransactionID"`
	PartyA                   int    `json:"PartyA"`
	IdentifierType           int    `json:"IdentifierType"`
	ResultURL                string `json:"ResultURL"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	Remarks                  string `json:"Remarks"`
	Occasion                 string `json:"Occasion"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
}
