package main

import (
	"fmt"
	"log"
	"net/http"

	"app-bazaar/backend"
	"app-bazaar/constants"
	"app-bazaar/handler"
)

func main() {
	if err := constants.LoadEnv(); err != nil {
        log.Println("No .env file found")
    }

	fmt.Println("Server started")

	backend.InitGCSBackend()
	backend.InitElasticBackend()

	log.Fatal(http.ListenAndServe(constants.SERVER_PORT, handler.InitRouter()))
}
