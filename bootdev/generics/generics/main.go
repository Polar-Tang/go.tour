package main

func getLast[T any](s []T) T {
	var zeroValue T
	if len(s) == 0 { // Handle empty slice
		return zeroValue
	}
	return s[len(s)-1]
}
