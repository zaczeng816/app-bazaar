package service

import (
	"errors"
	"fmt"
	"reflect"

	"app-bazaar/backend"
	"app-bazaar/constants"
	"app-bazaar/model"

	"github.com/olivere/elastic/v7"
)

func CheckUser(username, password string) (bool, error){
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("username", username))
	query.Must(elastic.NewTermQuery("password", password))

	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil{
		return false, err
	}

	var utype model.User
	for _, item := range searchResult.Each(reflect.TypeOf(utype)){
		u := item.(model.User)
		if u.Password == password{
			fmt.Printf("Logined in as %s\n", u.Username)
			return true, nil
		}
	}

	return false, errors.New("user not found")
}

func AddUser(user *model.User) (bool, error) {
	query := elastic.NewTermQuery("username", user.Username)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() > 0 {
		return false, errors.New("user already exists")
	}

	err = backend.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
	if err != nil {
		return false, err
	}
	
	fmt.Printf("User sucessfully created: %s\n", user.Username)
	return true, nil
}