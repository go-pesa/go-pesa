package main

import (
	"log"

	"github.com/go-pesa/internal/api"

	"github.com/go-pesa/internal/globals"
)

func init() {
	globals.InitConfig()
}

func main() {

	log.Printf("Yo World!")
	// x, _ := strconv.Atoi(globals.InitiatorShortCode)
	api.MakeB2CPayment(10, 254792561905, 602948, "BusinessPayment", "Hello test", "test")

}
