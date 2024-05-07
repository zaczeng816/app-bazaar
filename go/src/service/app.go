package service

import (
	"fmt"
	"reflect"

	"app-bazaar/backend"
	"app-bazaar/constants"
	"app-bazaar/model"

	"github.com/olivere/elastic/v7"
)

// SEARCH APPS

func SearchApps(title string, description string) ([]model.App, error){
	if title == "" {
		return SearchAppsByDescription(description)
	}
	if description == ""{
		return SearchAppsByTitle(title)
	}
	query1 := elastic.NewMatchQuery("title", title)
	query2 := elastic.NewMatchQuery("description", description)
	query := elastic.NewBoolQuery().Must(query1, query2)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.APP_INDEX)
	if err != nil{
		return nil, err
	}

	return getAppFromSearchResult(searchResult), nil
}


func SearchAppsByTitle(title string) ([]model.App, error){
	query := elastic.NewMatchQuery("title", title)
	query.Operator("AND")
	if title == ""{
		query.ZeroTermsQuery("all")
	}
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.APP_INDEX)
	if err != nil{
		return nil, err
	}

	return getAppFromSearchResult(searchResult), nil
}

func SearchAppsByDescription(description string) ([]model.App, error){
	query := elastic.NewMatchQuery("description", description)
	query.Operator("AND")
	if description == ""{
		query.ZeroTermsQuery("all")
	}
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.APP_INDEX)
	if err != nil{
		return nil, err
	}
	return getAppFromSearchResult(searchResult), nil
}

func SearchAppsById(appID string) (*model.App, error){
	query := elastic.NewTermQuery("id", appID)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.APP_INDEX)
	if err != nil{
		return nil, err
	}
	results := getAppFromSearchResult(searchResult)
	if len(results) == 1{
		return &results[0], nil
	}
	return nil, nil
}

func getAppFromSearchResult(searchResult *elastic.SearchResult) []model.App{
	var ptype model.App
	var apps []model.App
	for _, item := range searchResult.Each(reflect.TypeOf(ptype)){
		p := item.(model.App)
		apps = append(apps, p)
	}
	return apps
}

// SAVE

func SaveApp(app *model.App) error{
	productID, priceID, err := backend.CreateProductWithPrice(app.Title, app.Description, int64(app.Price*100))
	if err != nil{
		fmt.Printf("Failed to create product and price for app %v: %v\n", app.Title, err)
		return err
	}
	app.ProductID = productID
	app.PriceID = priceID
	fmt.Printf("Product and price created with: %v, %v\n", productID, priceID)
	return nil
}