package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	// Don't edit above this line

	var stringIssue string
	for _, issue := range issues {
		stringIssue += string(issue)
	}
	fmt.Println(stringIssue)
}
