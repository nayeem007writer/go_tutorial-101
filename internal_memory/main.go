package main

import "fmt"

var a = 10

func add(x, y int) {
	z := x + y
	fmt.Println("Sum is:", z)
}

func main() {
	fmt.Println("Hello, internal memory")
	add(5, 20)
	add(a, 15)
}

func init() {
	fmt.Println("Init function called")
}
