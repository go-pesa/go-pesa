package types

//OAuthString is a query string object for simple mapping
type OAuthString struct {
	GrantType string `url:"grant_type"`
}
