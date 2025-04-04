package main

import "strings"

func countDistinctWords(messages []string) int {

	counter := map[string]int{}

	for _, char := range messages {
		message := strings.ToLower(string(char))
		if _, ok := counter[strings.ToLower(message)]; !ok {
			counter[message]++
		}
	}
	return len(counter)
}
