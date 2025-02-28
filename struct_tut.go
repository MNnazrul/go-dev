package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p := Person("Name", 25)	
	fmt.Println(p)
}