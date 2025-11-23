package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func printDetails(usr User) {
	fmt.Println("User Details:")
	fmt.Println("Name:", usr.Name)
	fmt.Println("Email:", usr.Email)
	fmt.Println("Age:", usr.Age)
}

//receiver function only use for struct type
func (usr User) printDetailsUser() {
	fmt.Println("User Details:")
	fmt.Println("Name:", usr.Name)
	fmt.Println("Email:", usr.Email)
	fmt.Println("Age:", usr.Age)
}

func main() {

	var user1 User
	user1.Name = "Alice"
	user1.Email = "example@gmail.com"
	user1.Age = 30
	printDetails(user1)

	var user2 User
	user2.Name = "Bob"
	user2.Email = "	example2@gmail.com"
	user2.Age = 25
	user2.printDetailsUser()

}
