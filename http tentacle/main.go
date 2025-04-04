package main

import (
	"flag"
)

func main() {

	url := flag.String("urls", "", "URL to fetch")
	flag.Parse()

	fetchGet(url)

}
