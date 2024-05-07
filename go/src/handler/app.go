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
	service.SaveApp(&app)
	_, err := fmt.Fprintf(w, "Upload request processed: %s\n", app.Description)
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

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a checkout request")
	w.Header().Set("Content-Type", "text/plain")
	
	appID := r.FormValue("appID")
	s, err := service.CheckoutApp(r.Header.Get("Origin"), appID, )
	if err != nil{
		fmt.Println("Checkout failed")
		w.Write([]byte(err.Error()))
		return
	}	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s.URL))

	fmt.Println("Checkout process started!")
}