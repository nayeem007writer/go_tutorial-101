package main

import "fmt"


func addition(a int, b int)int {
	// This function adds two integers and prints the result.
	sum := a + b
	// fmt.Println("The sum is:", sum)
	return sum;
}

func subtraction(a int, b int)(int, int) {
	// This function subtracts two integers and returns the result.
	diff := a - b
	sum := a + b
	return diff, sum
}

func main() {
	// This is a placeholder for the main function.
	// You can add your application logic here.
	a:=5
	b:=10
	// sum := a + b
	p,q:=subtraction(a, b)
	fmt.Println("welcome to the world of go programing")
	var name string
	fmt.Println("Enter your name:_ _ _")
	fmt.Scanln(&name)
	fmt.Println("Hello", name, "the difference is:", p, "and the sum is:", q)
}