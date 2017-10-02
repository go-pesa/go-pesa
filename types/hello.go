package types

type Task struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	AppID     string `json:"appId"`
	Channel   string `json:"channel"`
	Site      string `json:"site"`
	RuleID    string `json:"ruleId"`
	Timestamp int64  `json:"timestamp"`
	MessageID string `json:"messageId"`
	Messages  []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ConnectionID string `json:"connectionId"`
		Timestamp    int64  `json:"timestamp"`
		Data         string `json:"data"`
	} `json:"messages"`
}
