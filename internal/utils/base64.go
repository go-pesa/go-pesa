package utils

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/go-pesa/internal/globals"
)

//EncodeConsumerKey creates a valid consumer key for fetching the access token
func EncodeConsumerKey() string {
	data := fmt.Sprintf("%s:%s", globals.ConsumerKey, globals.ConsumerSecret)
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	return uEnc
}
