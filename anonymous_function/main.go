package main

import "fmt"

func main() {
	// this is an anonymous function with IIfe (Immediately Invoked Function Expression)
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 6)
}
