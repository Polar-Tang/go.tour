package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetchGet(urls *string) {
	responseRecord := make(chan Response)
	defer close(responseRecord)

	file, err := os.Open(*urls)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	go func() {
		for dataChan := range responseRecord {
			fmt.Printf("Domain: %s, Status: %v, Protocol: %s\n",
				dataChan.Domain, dataChan.Status, dataChan.Protocol)
		}
	}()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}
		url := scanner.Text()
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
			}
			responseRecord <- Response{
				Domain:   url,
				Status:   res.StatusCode,
				Protocol: res.Proto,
				Length:   res.ContentLength,
			}
			defer res.Body.Close()
		}(url)

		select {
		case result := <-responseRecord:
			fmt.Println("Received response:", result)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout occurred")
		}

	}
}
