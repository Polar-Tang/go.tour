package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	if len(msgToCustomer) > 25 {
		return 0, fmt.Errorf("can't send texts over 25 characters")
	} else if len(msgToSpouse) > 25 {
		return 0, fmt.Errorf("can't send texts over 25 characters")
	}
	return (len(msgToCustomer) + len(msgToSpouse)) * 2, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
