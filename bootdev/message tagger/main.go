package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	messagestagged := make([]sms, len(messages))
	for key, msg := range messages {
		tags := tagger(msg)
		messagestagged[key] = sms{msg.id, msg.content, tags}
	}
	return messagestagged
}

func tagger(msg sms) []string {
	tags := []string{}
	urgent := "urgent"
	promo := "sale"

	isUrgent := strings.Contains(strings.ToLower(string(msg.content)), urgent)
	isPromo := strings.Contains(strings.ToLower(string(msg.content)), promo)

	if isUrgent && isPromo {
		return []string{"Urgent Promo"}
	}
	if isUrgent {
		return []string{"Urgent"}
	}
	if isPromo {
		return []string{"Promo"}
	}
	return tags
}
