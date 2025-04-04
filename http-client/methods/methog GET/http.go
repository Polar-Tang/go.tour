package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

func getUsers(url string) ([]User, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var users []User
	if err = json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

