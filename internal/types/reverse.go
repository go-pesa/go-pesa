package types

//ReverseTranactionRequest reverses a transaction
type ReverseTranactionRequest struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 int    `json:"Amount"`
	ReceiverParty          int    `json:"ReceiverParty"`
	RecieverIdentifierType int    `json:"RecieverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occasion               string `json:"Occasion"`
}
