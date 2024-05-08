package main

import (
	"fmt"
	"log"
	"net/http"

	"app-bazaar/backend"
	"app-bazaar/constants"
	"app-bazaar/handler"
	"app-bazaar/util"
)

func main() {
	conf, err := util.LoadApplicationConfig("conf", "deploy.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err = constants.LoadEnv(); err != nil {
        log.Println("No .env file found")
    }

	fmt.Println("Server started")

	backend.InitElasticBackend(conf.ElasticserachConfig)
	backend.InitGCSBackend(conf.GCSConfig)

	log.Fatal(http.ListenAndServe(constants.SERVER_PORT, handler.InitRouter()))
}
