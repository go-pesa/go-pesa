package utils

import (
	b64 "encoding/base64"
	"fmt"
)

//EncodeConsumerKey creates a valid consumer key for fetching the access token
func EncodeConsumerKey(ConsumerKey, ConsumerSecret string) string {
	data := fmt.Sprintf("%s:%s", ConsumerKey, ConsumerSecret)
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	return uEnc
}

//EncodePassword the payload with a ":" separator
func EncodePassword(shortCode, businessShortCode, timestamp string) string {
	data := fmt.Sprintf("%s%s%s", shortCode, businessShortCode, timestamp)
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	return uEnc
}
