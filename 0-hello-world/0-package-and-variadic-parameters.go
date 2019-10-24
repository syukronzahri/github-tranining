package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello everyone, this is the most interesting course on the earth")
	foo()
	fmt.Println("Something more...", n)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println(n)
	bar()
}

func foo() {
	fmt.Println("I'm in the foo")
}

func bar() {
	fmt.Println("And then we exited")
}
