package types

//IdentityRequest is an Identity Request Object
type IdentityRequest struct {
	Initiator         string `json:"Initiator"`
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	PhoneNumber       int    `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	TransactionDesc   string `json:"TransactionDesc"`
}
