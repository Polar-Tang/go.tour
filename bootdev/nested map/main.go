package main

func getNameCounts(names []string) map[rune]map[string]int {
	runemap := make(map[rune]map[string]int)

	for _, name := range names {
		if len(name) == 0 {
			continue
		}

		firstRune := rune(name[0])

		if _, exists := runemap[firstRune]; !exists {
			runemap[firstRune] = make(map[string]int)
		}

		runemap[firstRune][name]++
	}

	return runemap
}
