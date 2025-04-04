package main

import "strings"

func removeProfanity(message *string) {
	*message = strings.Replace(*message, "fubb", "****", 2)
	*message = strings.Replace(*message, "shiz", "****", 2)
	*message = strings.Replace(*message, "witch", "*****", 2)

}
