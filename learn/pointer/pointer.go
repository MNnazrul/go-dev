package main

import "fmt"

type User struct {
	ID string
	Username string
}

type Cart struct {
	ID string
	Paid bool
}

func main() {

	user1 := User{"123", "name"}
	var userPtr *User = &user1 

	intVal := 12
	var intPtr *int = &intVal

	fmt.Println(intPtr) 
	fmt.Println(*intPtr) // Dereferencing

	fmt.Println(userPtr)

	cart := Cart{
		ID: "12123",
		Paid: true,
	}

	cartPtr := &cart

	cartDeref := *cartPtr

	fmt.Println(cartPtr)
	fmt.Println(cartDeref)

}