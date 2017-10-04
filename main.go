package main

import (
	"log"
	"strconv"

	"github.com/go-pesa/internal/api"

	"github.com/go-pesa/internal/globals"
)

func init() {
	globals.InitConfig()
}

func main() {

	log.Printf("Yo World!")
	x, _ := strconv.Atoi(globals.InitiatorShortCode)
	api.MakeB2BPayment(10, 600000, 4, x, 4, "BusinessToBusinessTransfer", "xcv12de", "sadasd asc")

}
