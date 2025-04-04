package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	hunterUrl = "https://api.hunter.io/v2/domain-search"
	domain    = "stripe.com"
	apiKey    = "7ebc451431d0616a9c542312242a8c4c8107aa64" // Replace with your actual Hunter API key
)

type HunterResponse struct {
	Data struct {
		Domain string `json:"domain"`
		// ... other fields
		Emails []Email `json:"emails"`
	} `json:"data"`
}

type Email struct {
	Value      string `json:"value"`
	Type       string `json:"type"`
	Confidence int    `json:"confidence"`
	// ... other fields
}

func main() {

	file, err := os.Open("firebounty_programs1.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		contentLine := scanner.Text()
		words := strings.Fields(contentLine)
		for _, word := range words {

			url := fmt.Sprintf("%s?domain=%s&api_key=%s", hunterUrl, word, apiKey)

			fmt.Printf("Scanning: %s\n", url)

			client := &http.Client{}

			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// -----------------------------------------------------------
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("Error:", resp.Status)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}
			// ------------------------------------------------

			var hunterResponse HunterResponse
			err = json.Unmarshal(body, &hunterResponse)
			if err != nil {
				fmt.Println("Error unmarshaling response:", err)
				return
			}
			emails := hunterResponse.Data.Emails

			emailSlice := make([]string, len(emails))
			for _, email := range emails {
				emailSlice = append(emailSlice, email.Value)
			}

			theMap := map[string][]string{}

			fmt.Println(theMap)
			theMap[word] = emailSlice
		}
	}

}
