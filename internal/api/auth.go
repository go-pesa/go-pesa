package api

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-pesa/internal/globals"

	"github.com/go-pesa/internal/http"
	"github.com/go-pesa/internal/types"
	"github.com/go-pesa/internal/utils"
)

//generateToken generates an AccessToken
func generateToken() string {

	endpoint := "oauth/v1/generate"

	var headers = []types.Header{
		types.Header{
			Key:   "Authorization",
			Value: fmt.Sprintf("Basic %s", utils.EncodeConsumerKey()),
		},
	}

	params := &types.OAuthString{
		GrantType: "client_credentials",
	}

	res := http.Get(endpoint, headers, params)

	var token types.AccessToken

	utils.Unmarshal(res, &token)
	cacheToken(token)
	return token.AccessToken

}

//GetToken either retrives token from Cache or Get a new one
func GetToken() string {
	token, err := fetchToken()
	if err != nil {
		return generateToken()
	}
	return fmt.Sprintf("%s", token)

}

//cacheToken caches token with the expiry time received from API
func cacheToken(token types.AccessToken) {
	expiry, err := strconv.Atoi(token.ExpiresIn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error converting expiry to string, error:%s", err))
	}
	globals.Cache.SetWithExpire("Token", token.AccessToken, time.Duration(expiry-100)*time.Second)
}

//fetches stored token if not expired
func fetchToken() (interface{}, error) {
	value, err := globals.Cache.Get("Token")
	return value, err
}
