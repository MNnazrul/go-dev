package main

import (
	"fmt"
	"io"
)

type Human struct {
	Firstname string
	Lastname string
	Age int
	Country string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GivenAffection(to Human)
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s", c.Name, from.Firstname)
}

func (c Cat) GivenAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human %s", c.Name, to.Firstname)
}

func (d Dog) ReceiveAffection(from Human) {
    fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
    fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GivenAffection(human)
	animal.ReceiveAffection(human)
}

func main() {
	var r io.Reader
	fmt.Println(r)
}