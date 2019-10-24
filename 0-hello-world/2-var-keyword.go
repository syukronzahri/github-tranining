package main

import "fmt"

// declare the variable "y"
// assign the value 43
// declare & assign = initialization
var y = 43

// declare there is a variable with the identifier "z"
// and that the variable with the identifier "z" is of TYPE int
// assigns the zero value of type int to "z"
//  false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// declare a variable and assign a value (of certain type)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
