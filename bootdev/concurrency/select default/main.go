package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	select {
	case <-snapshotTicker:
		time.After(2000 * time.Millisecond)
		go takeSnapshot(logChan)
	case <-saveAfter:
		go saveSnapshot(logChan)
	default:
		waitForData(logChan)
		time.Sleep(500 * time.Millisecond)
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
