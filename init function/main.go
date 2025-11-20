package main

import "fmt"

var a = 90

func main() {
	fmt.Println("Hello, init function ")
	fmt.Println("Value of a in main function is:", a)
}

func init() {
	fmt.Println("Init function called and value of a is:", a)
	a = 80
}

//here init function will be called first before main function
