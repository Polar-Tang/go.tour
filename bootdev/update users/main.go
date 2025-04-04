package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	if membershipType == "premium" {
		u := User{
			Name: name,
			Membership: Membership{
				MessageCharLimit: 1000,
				Type:             "premium",
			},
		}
		return u

	}
	if membershipType == "standard" {
		u := User{
			Name: name,
			Membership: Membership{
				MessageCharLimit: 100,
				Type:             "standard",
			},
		}
		return u

	}
	return User{}
}

func main() {
	fmt.Println(reflect.TypeOf(User{}.Membership))
}
