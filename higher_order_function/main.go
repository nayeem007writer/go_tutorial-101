package main

import "fmt"

/*
take a function as parameter
return a function from another function
and both are called higher order function
*/
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func multiply(a, b int) int {
	return a * b
}

func main() {

	sum := operate(5, 20, multiply)
	fmt.Println("Sum is:", sum)
	// passing function as parameter
}
