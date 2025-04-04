package main

func countReports(numSentCh chan int) int {
	var v int

	for {
		val, ok := <-numSentCh
		if !ok {
			break
		}
		v = val + v
	}

	return v

}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
