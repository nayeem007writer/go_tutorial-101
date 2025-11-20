package main

import "fmt"

func a(a, b int) {
	fmt.Println(a + b)
}

func main() {
	a(3, 4)
	a := func(a, b int) {
		fmt.Println(a + b)
	}
	a(7, 8)
}
