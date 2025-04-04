package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	var issues []Issue

	if err := json.Unmarshal(data, &issues); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return issues, nil
}
