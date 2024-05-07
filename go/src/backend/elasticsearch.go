package backend

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"

	"app-bazaar/constants"
)

var (
	ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct{
	client *elastic.Client
}

func InitElasticBackend(){
	client, err := elastic.NewClient(
		elastic.SetURL(constants.ES_URL),
		elastic.SetBasicAuth(constants.ES_USERNAME, constants.ES_PASSWORD))

	if err != nil{
		panic(err)
	}

	exists, err := client.IndexExists(constants.APP_INDEX).Do(context.Background())
	if err != nil{
		panic(err)
	}

	if !exists{
		mapping := `{
			"mappings": {
				"properties": {
					"id": {"type": "keyword"},
					"user": {"type": "keyword"},

					"title": {"type": "text"},
					"description": {"type": "text"},

					"price": {"type": "keyword", "index": false},
					"url": {"type": "keyword", "index": false}
				}
			}
		}`
		
		_, err := client.CreateIndex(constants.APP_INDEX).Body(mapping).Do(context.Background())

		if err != nil{
			panic(err)
		}

	}

	exists, err = client.IndexExists(constants.USER_INDEX).Do(context.Background())
	if err != nil{
		panic(err)
	}

	if !exists{
		mapping := `{
			"mappings": {
				"properties": {
					"username": {"type": "keyword"},
					"password": {"type": "keyword"},
					"age": {"type": "long", "index": false},
					"gender": {"type": "keyword", "index": false}
				}
			}
		}`
		
		_, err := client.CreateIndex(constants.USER_INDEX).Body(mapping).Do(context.Background())

		if err != nil{
			panic(err)
		}

	}

	fmt.Println("Indices created")
	ESBackend = &ElasticsearchBackend{client: client}

}

func (es *ElasticsearchBackend) ReadFromES(query elastic.Query, index string) (*elastic.SearchResult, error){
	searchResult, err := es.client.Search().
		Index(index).
		Query(query).
		Do(context.Background())

	if err != nil{
		return nil, err
	}

	return searchResult, nil

}