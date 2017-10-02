package main

import (
	"log"

	"github.com/go-pesa/internal/utils"

	"github.com/go-pesa/internal/globals"
)

func init() {
	globals.InitConfig()
}

func main() {
	log.Printf(utils.EncodeConsumerKey())
}
