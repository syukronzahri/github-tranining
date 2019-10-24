package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most interesting course on the earth")
	foo()
	fmt.Println("Something more...")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in the foo")
}

func bar() {
	fmt.Println("And then we exited")
}
