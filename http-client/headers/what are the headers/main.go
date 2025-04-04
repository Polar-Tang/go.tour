package main

import (
	"net/http"
)

func getContentType(res *http.Response) string {

	resp := res.Header.Get("Content-Type")

	return resp
}
