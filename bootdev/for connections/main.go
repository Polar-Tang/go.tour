package main

func countConnections(groupSize int) int {
	var acc int
	for i := 0; i < groupSize; i++ {
		acc += groupSize - i - 1
	}
	return acc
}
