package main

import (
	"log"

	"github.com/go-pesa/internal/globals"
)

func init() {
	globals.InitConfig()
}

func main() {

	log.Printf("Yo World!")
	// api.RequestBalance(globals.InitiatorShortCode, 4, "hello")

}
