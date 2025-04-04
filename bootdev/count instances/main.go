package main

func getCounts(messagedUsers []string, validUsers map[string]int) {

	for _, char := range messagedUsers {
		message := string(char)
		if _, ok := validUsers[message]; ok {
			validUsers[message]++
		}
		validUsers = map[string]int{}
	}
}
