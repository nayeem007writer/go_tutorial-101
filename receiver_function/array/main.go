package main

import "fmt"

func main() {
	var array1 [2]int

	//short hand declaration of array
	arr := [3]string{"Go", "Python", "Java"}
	arr2 := [3]int{10, 20, 30}

	array1[0] = 10
	array1[1] = 20

	fmt.Println("array elements", array1)
	fmt.Println("array elements", arr)
	fmt.Println("array elements", arr2)

}
