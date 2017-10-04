package types

//GenericResponse is the response object
type GenericResponse struct {
	OriginatorConverstionID string `json:"OriginatorConverstionID"`
	ConversationID          string `json:"ConversationID"`
	ResponseDescription     string `json:"ResponseDescription"`
}
