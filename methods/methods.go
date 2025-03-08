package main

import (
	"fmt"
	"time"
)

type Teacher struct {
	id int
	firstname string
	lastname string
}

func (t *Teacher) sayHelloToClass() {
	fmt.Printf("Hi class, my name is %s %s\n", t.firstname, t.lastname)
	t.firstname = "Mick"
}

func (t Teacher) sayHelloToClass1() {
	fmt.Printf("Hi class, my name is %s %s\n", t.firstname, t.lastname)
	t.firstname = "Mick"
}

func main() {
	fmt.Println(time.Now())
	teacher := Teacher{12, "John", "Doe"}
	teacher.sayHelloToClass1()
	teacher.sayHelloToClass1()
	teacher.sayHelloToClass()
	teacher.sayHelloToClass()
}