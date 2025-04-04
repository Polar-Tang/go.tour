package main

import (
	"fmt"
	"net/http"
)

func fetch() (int, error) {
	res, err := http.Get("https://buckets.grayhatwarfare.com/files?keywords=-png+-jpg+-jpeg+-svg+-webp+-gif+-js+-css+-scss+-static+-jfif+-avif&order=size&direction=desc&bucket=ipimages.s3-us-west-2.amazonaws.com")
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	res.Header.Add("Authorization", " Bearer 54e7fe8c2aa1dd504b9be39fa3466f10")

	return res.StatusCode, nil
}
