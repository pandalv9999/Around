package main

import (
	"fmt"
	"reflect"
	"github.com/olivere/elastic"
)

const UserIndex = "user"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age int64 `json:"age"`
	Gender string `json:"gender"`
}

func checkUser(username, password string) (bool, error) {
	query := elastic.NewTermQuery("username", username)
	searchResult, err := readFromElasticSearch(query, UserIndex)
	if err != nil {
		return false, err
	}
	var utype User
	for _, item := range searchResult.Each(reflect.TypeOf(utype)){
		if u, ok := item.(User); ok {
			if u.Password == password {
				fmt.Printf("Login as %s\n", username)
				return true, nil
			}
		}
	}
	return false, nil
}

func addUser(user *User) (bool, error) {
	query := elastic.NewTermQuery("username", user.Username)
	searchResult, err := readFromElasticSearch(query, UserIndex)
	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() > 0 {
		return false, nil
	}

	err = saveToElasticSearch(user, UserIndex, user.Username)
	if err != nil {
		return false, err
	}

	fmt.Printf("User is added: %s\n", user.Username)
	return true, nil
}