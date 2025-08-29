package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 6)
	add(p, a)
}

func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}

/*

--

Variable shadowing occurs when a variable declared within a certain scope (like a function or a block) has the same name as a variable in an outer scope.

The inner variable effectively “shadows” the outer variable, meaning that the inner variable is accessible in that scope while the outer variable is temporarily hidden or inaccessible.


==============

how go code run: 2 phases
1. compile. go build main.go
2. execution. ./main


****** compile phase ********

code segment: const variable and all the func defination.

	a = 10
	call = func ()  {...}
	add = func () {...}
	main = func () {...}
	init = func() {...}

* stack frame.

data segment => global variable.

===============



===============

*/
