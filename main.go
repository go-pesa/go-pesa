package main

import (
	"fmt"
	"log"

	"github.com/go-pesa/internal/api"

	"github.com/go-pesa/internal/globals"
)

func init() {
	globals.InitConfig()
}

func main() {

	log.Printf(fmt.Sprintf("%s", api.GetToken()))

	log.Printf(fmt.Sprintf("%s", api.GetToken()))

}
