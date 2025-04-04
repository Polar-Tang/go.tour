package main

func concurrentFib(n int) []int {
	// create a channel of ints
	ch := make(chan int)
	var sliceToReturn []int

	// call fibonacci concurrently
	fibonacci(n, ch)

	// Use a range loop to read from the channel and append the values to a slice
	for item := range ch {
		sliceToReturn = append(sliceToReturn, item)
	}

	return sliceToReturn
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
