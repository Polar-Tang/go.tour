package main

func bulkSend(numMessages int) float64 {
	var acc float64
	for i := 0; i < numMessages; i++ {
		acc += 1 + (float64(i) / 100)
	}
	return acc
}
