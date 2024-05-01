package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app-bazaar/model"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a upload request")
	decoder := json.NewDecoder(r.Body)
	var app model.App
	if err := decoder.Decode(&app); err != nil {
		panic(err)
	}
	_, err := fmt.Fprintf(w, "Upload request received: %s\n", app.Description)
	if err != nil {
		return
	}
}
