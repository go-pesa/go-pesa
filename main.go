package main

import (
	"log"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/types"
)

func init() {
	globals.InitConfig()
}

func main() {
	var task types.Task

	task.ID = "12"

	log.Printf(globals.ConsumerSecret)
}
