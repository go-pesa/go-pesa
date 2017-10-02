package globals

import (
	"fmt"
	"log"
	"os"

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
	sandboxURL = "https://production.safaricom.co.ke/"
	//productionURL is the mpesa api production URL
	productionURL = "https://production.safaricom.co.ke/" //TO-DO: Update with correct URL
)

//InitConfig initializes global variables and all other environment variables
func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Using os environment variables")
	} else {
		log.Println("Using .env for configuration")
	}

	BaseURL = setEnvironment(os.Getenv("APP_ENV"))
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
