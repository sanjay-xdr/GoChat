package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Run this file")

	mux := routes()

	log.Println("Server is Starting at index 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
