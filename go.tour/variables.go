// Every script in Go is made up of packages.
package main

// Programs start running in the package main

/*
The import could be for each package
import "fmt"
import "math"
*/
// but this way it's better
import (
	"fmt"
	"math"
	/*// by convention the package name it's the last name in the path
	"math/rand"
	*/)

/*
The args could be typed for each
x int, y int
*/
// when two args share the type, we use a single type by the right
func add(x, y int) int {
	return x + y
}

/*
A function may return many variables:
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
A named function specify the return values at the function signature, then use a single return at the bottom of the gunction body
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool // false false false

// Constants can be character, string, boolean, or numeric values.
const Pi = 3.14

// Constants cannot be declared using the := syntax.

// constant may change of value
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func AllVars() {

	// VARIABLES ----------------------------------------------------
	// fmt.Println("My favorite number is", rand.Intn(10))

	// the variables which start with a capital letter are the exported variables. Pi it's exported from math, that's why it should start with a capital letter
	fmt.Println(math.Pi)

	var i int
	// var initialized
	var c, python, java = true, false, false
	fmt.Println(i, c, python, java)

	// walrus = declare and intialize, initializes are not allowed outside a function
	// php := false

	// TYPES ----------------------------------------------------

	/*
		GO BASIC TYPES
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		     // represents a Unicode code point

		float32 float64

		complex64 complex128
	*/

	// The expression T(v) converts the value v to the type T.
	f := float64(i)
	u := uint(f)

	fmt.Printf("u is of type %T\n", u)

}
