package api

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bluele/gcache"
	"github.com/go-pesa/go-pesa/internal/http"
	"github.com/go-pesa/go-pesa/internal/types"
	"github.com/go-pesa/go-pesa/internal/utils"
)

//generateToken generates an AccessToken
func generateToken(key, secret, baseURL string, cache gcache.Cache) string {

	endpoint := oauthGenerate

	var headers = []types.Header{
		types.Header{
			Key:   "Authorization",
			Value: fmt.Sprintf("Basic %s", utils.EncodeConsumerKey(key, secret)),
		},
	}

	params := &types.OAuthString{
		GrantType: "client_credentials",
	}

	res := http.Get(endpoint, baseURL, headers, params)

	var token types.AccessToken

	utils.Unmarshal(res, &token)
	cacheToken(token, cache)
	return token.AccessToken

}

//getToken either retrives token from Cache or Get a new one
func getToken(key, secret, baseURL string, cache gcache.Cache) string {
	token, err := fetchToken(cache)
	if err != nil {
		return generateToken(key, secret, baseURL, cache)
	}
	return fmt.Sprintf("%s", token)
}

//cacheToken caches token with the expiry time received from API
func cacheToken(token types.AccessToken, cache gcache.Cache) {
	expiry, err := strconv.Atoi(token.ExpiresIn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error converting expiry to string, error:%s", err))
	}
	cache.SetWithExpire("Token", token.AccessToken, time.Duration(expiry-100)*time.Second)
}

//fetches stored token if not expired
func fetchToken(cache gcache.Cache) (interface{}, error) {
	value, err := cache.Get("Token")
	return value, err
}
