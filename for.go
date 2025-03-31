// the for it's the only iterable in go
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// ------------------------------------- FOR

func initialStatement() {
	// the initial stamet is block scoped
	// it can be omitted
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func whileInGo() {
	// or also the whole for could by the initial statement, which creates a while loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
		fmt.Println("DIE")
	}
}
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i > 20; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return float64(z)
}

// ------------------------------------- IF

func sqrt(x float64) string {
	// IF also aren't surronded by parenthesis
	// handle imagine numbers
	if x < 0 {
		return sqrt(x) + "i"
	}
	// return the square
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// the if could declares variables in block
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// ------------------------------------- SWITCH

func switchCases() {
	// switch evaluates brom the top to bottom and breaks on the first succeded condition.

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func whenSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func AllFor() {
	initialStatement()

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	switchCases()
}
