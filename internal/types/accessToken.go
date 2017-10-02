package types

//AccessToken object received from successfully performing client credential grant type
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}
