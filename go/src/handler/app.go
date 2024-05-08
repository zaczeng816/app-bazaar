package handler

import (
	"app-bazaar/model"
	"app-bazaar/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pborman/uuid"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a upload request")
	// decoder := json.NewDecoder(r.Body)
	// var app model.App
	// if err := decoder.Decode(&app); err != nil {
	// 	panic(err)
	// }
	app := model.App{
		Id: uuid.New(),
		User: r.FormValue("user"),
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	app.Price = price
	fmt.Printf("%v, %T", price, price)
	if err != nil {
		fmt.Println(err)
	}
	file, _, err := r.FormFile("media_file")
	if err != nil {
		http.Error(w, "Media file is not available", http.StatusBadRequest)
		fmt.Printf("Media file is not available %v\n", err)	
		return
	} 

	err = service.SaveApp(&app, file)
	if err != nil {
		http.Error(w, "Failed to save app to backend", http.StatusInternalServerError)
		fmt.Printf("Failed to save app to backend %v\n", err)	
		return
	}
	fmt.Println("App successfully saved to backend")
	
	fmt.Fprintf(w, "App was saved successfully %s\n", app.Description)
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