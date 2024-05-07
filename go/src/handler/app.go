package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app-bazaar/model"
	"app-bazaar/service"
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

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a search request")
	w.Header().Set("Content-Type", "application/json")
	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")

	var apps []model.App 
	var err error
	apps, err = service.SearchApps(title, description)
	if err != nil{
		http.Error(w, "Fail to read Apps from backend", http.StatusInternalServerError)
		return 
	}
	
	js, err := json.Marshal(apps)
	if err != nil{
		http.Error(w, "Failed to fail Apps into JSON", http.StatusInternalServerError)
	}
	w.Write(js)
}


