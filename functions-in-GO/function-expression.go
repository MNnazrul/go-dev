package main

import "fmt"

var sum = func() {
	add(3, 4)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	sum()
	add(4, 4)
	add := func(a , b int) {
		fmt.Println(a + b)
	}
	add(5, 6)
}