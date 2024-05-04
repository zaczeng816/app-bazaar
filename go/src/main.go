package main

import (
	"fmt"
	"log"
	"net/http"

	"app-bazaar/backend"
	"app-bazaar/handler"
)

func main() {
	fmt.Println("Server started")

	backend.InitElasticBackend()

	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
