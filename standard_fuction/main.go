package main

import "fmt"

// this is a function with named is called named function or standard function
func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	add(3, 4)
}
