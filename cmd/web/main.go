package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjay-xdr/ws/internals/handlers"
)

func main() {
	fmt.Print("Run this file")

	mux := routes()

	go handlers.ListenToWsChannel()
	log.Println("Server is Starting at index 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
