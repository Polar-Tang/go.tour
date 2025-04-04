package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 0; i < 40; i++ {
		if i == 2 {
			fmt.Println(i)
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
