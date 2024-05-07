package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"app-bazaar/backend"
	"app-bazaar/constants"
	"app-bazaar/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

	fmt.Println("Server started")

	backend.InitElasticBackend()

	log.Fatal(http.ListenAndServe(constants.SERVER_PORT, handler.InitRouter()))
}
