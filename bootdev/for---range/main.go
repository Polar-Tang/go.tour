package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for _, word := range msg {
		// if word == badWords[word]{
		// 	return badWords[word]
		// }
		for key, badWord := range badWords {
			if badWord == word {
				return key + 1
			}
		}

	}
	return -1
}
