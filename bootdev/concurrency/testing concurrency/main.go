package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // initialize it
	c2 := make(chan string) // initialize it

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
