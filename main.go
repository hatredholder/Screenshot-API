package main

import (
	"log"

	"github.com/hatredholder/Screenshot-API/api"
)

func main() {
	server := api.NewServer(":8080")

	err := server.Start()
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Printf("Error occured when starting the server: %s", err)
	}
}
