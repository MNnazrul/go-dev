package main

import "fmt"

// Maps and channels are already references to the internal structure. Consequently, a function/method that accepts a map or a channel can modify it, even if the parameter is not a pointer type.


func addElement(cities map[string]string) {
	cities["france"] = "paris"
}

func main() {

	cities := make(map[string]string)
	addElement(cities)
	fmt.Println(cities)
	
}