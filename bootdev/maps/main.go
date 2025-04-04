package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	newUser := make(map[string]user)

	for i := 0; i < len(names); i++ {
		newUser[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return newUser, nil
}

type user struct {
	name        string
	phoneNumber int
}
