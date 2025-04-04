package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * len(e.body)
	} else if !e.isSubscribed {
		return 5 * len(e.body)
	}
	return 0
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | %s", e.body, "Subscribed")
	} else if !e.isSubscribed {
		return fmt.Sprintf("'%s' | %s", e.body, "Not Subscribed")
	}
	return ""
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
