package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return data, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
