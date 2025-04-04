package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {

	req, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	if req.StatusCode != http.StatusOK {
		return req.StatusCode, fmt.Errorf("non-OK HTTP status: %s", req.Status)
	}

	defer req.Body.Close()

	return req.StatusCode, nil
}
