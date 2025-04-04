package main

import "net/http"

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
