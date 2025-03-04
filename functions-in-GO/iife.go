package main

import "fmt"

func main() {
	func(a int, b int) {
		fmt.Println(a + b)
	} (3, 5)

}

func init() {
	fmt.Println("I'll be called first")
}