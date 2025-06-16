package main

import (
	"fmt"
	"io"
)

/*
./interface-embedding.go:68:6: cannot use c (variable of type Cat) as DomesticAnimal value in argument to Pet: Cat does not implement DomesticAnimal (missing method String)
./interface-embedding.go:69:6: cannot use d (variable of type Dog) as DomesticAnimal value in argument to Pet: Dog does not implement DomesticAnimal (missing method String)
*/

type Human struct {
	Firstname string
	Lastname string
	Age int
	Country string
}

type Stringer interface {
	String() string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GivenAffection(to Human)
	Stringer
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

func (c Cat) String() string {
	return c.Name
}

func (d Dog) String() string {
	return d.Name
}

func (d Dog) ReceiveAffection(from Human) {
    fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GivenAffection(to Human) {
    fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GivenAffection(human)
	animal.ReceiveAffection(human)
}

func main() {
	var r io.Reader
	fmt.Println(r)

	var john Human
	john.Firstname = "John"

	var c Cat
	c.Name = "Maru"

	var d Dog
	d.Name = "Medor"


	Pet(c, john)
	Pet(d, john)

}