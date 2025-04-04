package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if user.scheduledForDeletion {
		return true, nil
	}
	return false, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
