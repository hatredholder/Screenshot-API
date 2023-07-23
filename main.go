package main

import (
	"fmt"
	"log"

	"github.com/hatredholder/Screenshot-API/api"
)

func main() {
	server := api.NewServer(":8080")

	fmt.Printf("[INFO] Server started on port: \"%s\"\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.SetPrefix("[ERROR]")
		log.Printf("Error occured when starting the server: %s", err)
	}
}
