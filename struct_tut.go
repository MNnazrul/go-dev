package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	var user1 User

	user1 = User {
		Name : "Nazrul",
		Age : 25,
	}

	fmt.Println(user1)

}