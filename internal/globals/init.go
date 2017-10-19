package globals

import (
	"github.com/bluele/gcache"
)

var (
	//Cache to hold temporary info like auth tokens
	Cache gcache.Cache
	//BaseURL holds the URL being used
	BaseURL string
)

//InitConfig initializes global variables and all other environment variables
func InitConfig(key, secret, url string) {

	Cache, BaseURL = gcache.
		New(20).
		LRU().
		Build(), url
	Cache.Set("CONSUMER_KEY", key)
	Cache.Set("CONSUMER_SECRET", secret)

}

// func makeSecurityCredential() string {
// 	pemData := certs.MpesaCert
// 	block, _ := pem.Decode(pemData)
// 	if block == nil {
// 		log.Fatalf("bad key data: %s", "not PEM-encoded")
// 	}
// 	if got, want := block.Type, "CERTIFICATE"; got != want {
// 		log.Fatalf("unknown key type %q, want %q", got, want)
// 	}
// 	key, err := x509.ParseCertificate(block.Bytes)
// 	if err != nil {
// 		log.Fatalf("bad private key: %s", err)
// 	}
// 	pkey := key.PublicKey.(*rsa.PublicKey)
// 	out, err := rsa.EncryptPKCS1v15(rand.Reader, pkey, []byte(mpesaSecret))
// 	if err != nil {
// 		log.Fatalf("decrypt: %s", err)
// 	}
// 	return base64.URLEncoding.EncodeToString(out)

// }
