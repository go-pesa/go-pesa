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
	// api.ReverseTransaction("LJ38IOUAHY",)
	// api.StkPushProcess(1, 254792561905, "test", "test")
	api.ReverseTransaction("LJ40JBB8N8", 1, 602948, 4, "testtes", "test")
}
