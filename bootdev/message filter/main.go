package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var filteredMessages []Message
	for _, message := range messages {
		switch filterType {
		case "text":
			if filterType == "text" {
				filteredMessages = append(filteredMessages, message)
			}
			continue
		case "media":
			if filterType == "media" {
				filteredMessages = append(filteredMessages, message)
			}
			continue

		case "link":
			if filterType == "link" {
				filteredMessages = append(filteredMessages, message)
			}
			continue

		case "Check":
			if filterType == "Check" {
				filteredMessages = append(filteredMessages, message)
			}
			continue

		}

	}
	return filteredMessages

}
