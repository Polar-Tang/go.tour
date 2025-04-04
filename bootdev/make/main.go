package main

func getMessageCosts(messages []string) []float64 {
	messageLen := len(messages)

	costSlice := make([]float64, messageLen)
	for i := 0; i < messageLen; i++ {
		costSlice[i] = float64(len(messages[i])) * 0.01
	}
	return costSlice
}
