package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	var users []User
	var fullURL string
	for _, user := range users {
		fullURL = url + "?sort=" + user.User.Name
	}
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
