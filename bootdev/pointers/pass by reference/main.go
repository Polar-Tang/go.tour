package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(testing *Analytics, message Message) Analytics {

	if message.Success {
		testing.MessagesSucceeded++
	}
	if !message.Success {
		testing.MessagesFailed++
	}
	testing.MessagesTotal++

	return *testing
}
