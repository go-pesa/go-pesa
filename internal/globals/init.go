package globals

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/bluele/gcache"
	"github.com/go-pesa/internal/certs"
	"github.com/joho/godotenv"
)

var (
	//BaseURL is the base URl for all endpoints, this depends on whether it is a dev or production environment
	BaseURL string
	//ConsumerSecret is the mpesa consumer secret, please set it with envirnment variables as MPESA_CONSUMER_SECRET
	ConsumerSecret string
	//ConsumerKey is the mpesa consumer key, please set it with envirnment variables as MPESA_CONSUMER_KEY
	ConsumerKey string
	//sandboxURL is the mpesa sandbox url
	sandboxURL = "https://sandbox.safaricom.co.ke/"
	//productionURL is the mpesa api production URL
	productionURL = "https://production.safaricom.co.ke/" //TO-DO: Update with correct URL
	//Cache to hold temporary info like auth tokens
	Cache gcache.Cache
	//ShortCode from env
	ShortCode string
	//PassKey from env
	PassKey string
	//Callback url for mpesa transactions
	Callback string
	//MSISDN is the MSISDN
	MSISDN string
	//Initiator as defined by the docmentation
	Initiator string
	//InitiatorShortCode is the shortcode for initiator transactions
	InitiatorShortCode string
	//MpesaSecret
	mpesaSecret string
	//SecurityCredential encrypted MpesaSecret
	SecurityCredential string
)

//InitConfig initializes global variables and all other environment variables
func InitConfig() {

	Cache = gcache.
		New(20).
		LRU().
		Build()

	_ = godotenv.Load()

	BaseURL = setEnvironment(os.Getenv("APP_ENV"))
	setVariable(os.Getenv("MPESA_TEST_MSISDN"), &MSISDN, "M-pesa MSISDN")
	setVariable(os.Getenv("MPESA_CALLBACK_URL"), &Callback, "M-pesa Callback")
	setVariable(os.Getenv("LIPA_NA_MPESA_PASSKEY"), &PassKey, "M-pesa PassKey")
	setVariable(os.Getenv("LIPA_NA_MPESA_SHORTCODE"), &ShortCode, "M-pesa ShortCode")
	setVariable(os.Getenv("MPESA_CONSUMER_KEY"), &ConsumerKey, "M-pesa Consumer Key")
	setVariable(os.Getenv("MPESA_CONSUMER_SECRET"), &ConsumerSecret, "M-pesa Consumer Secret")
	setVariable(os.Getenv("MPESA_INITIATOR_CODE_1"), &Initiator, "M-pesa initiator")
	setVariable(os.Getenv("MPESA_SHORTCODE_1"), &InitiatorShortCode, "M-pesa initiator shortcode")
	setVariable(os.Getenv("MPESA_SECRET"), &mpesaSecret, "M-pesa secret for security credentials")
	SecurityCredential = makeSecurityCredential()

}

func setEnvironment(env string) string {
	if env == "production" {
		return productionURL
	}
	return sandboxURL
}

func setVariable(value string, key *string, variable string) {
	if value == "" {
		log.Fatal(fmt.Sprintf("%s not set", variable))
	}
	*key = value
}

func makeSecurityCredential() string {
	pemData := certs.MpesaCert
	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Fatalf("bad key data: %s", "not PEM-encoded")
	}
	if got, want := block.Type, "CERTIFICATE"; got != want {
		log.Fatalf("unknown key type %q, want %q", got, want)
	}
	key, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("bad private key: %s", err)
	}
	pkey := key.PublicKey.(*rsa.PublicKey)
	out, err := rsa.EncryptPKCS1v15(rand.Reader, pkey, []byte(mpesaSecret))
	if err != nil {
		log.Fatalf("decrypt: %s", err)
	}
	return string(out)
}
