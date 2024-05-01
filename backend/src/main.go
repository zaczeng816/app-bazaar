package main

import (
	"fmt"
	"log"
	"net/http"

	"app-bazaar/handler"
)

func main() {
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
