package globals

import (
	"fmt"
	"log"
	"os"

	"github.com/bluele/gcache"
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
