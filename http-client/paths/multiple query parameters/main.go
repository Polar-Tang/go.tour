package main

import "strconv"

func fetchTasks(baseURL, availability string) []Issue {
	var issue Issue
	var availabilityRank string

	if availability == "Low" {
		availabilityRank = "1"
	} else if availability == "Medium" {
		availabilityRank = "2"
	} else if availability == "High" {
		availabilityRank = "3"
	}

	fullURL := baseURL + "?sort=" + strconv.Itoa(issue.Estimate) + "&limit=" + availabilityRank

	return getIssues(fullURL)
}
