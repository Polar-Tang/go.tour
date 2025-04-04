package main

import "fmt"

func fizzbuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz ", i)
		}
		if i%5 == 0 {
			fmt.Println("Buzz ", i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
