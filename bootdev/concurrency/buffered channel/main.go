package main

func addEmailsToQueue(emails []string) chan string {
	c := make(chan string, len(emails))

	for _, email := range emails {
		c <- email

	}
	return c
}
