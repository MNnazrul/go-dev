package main

import (
	"fmt"
	"os"
)

func main() {
	
	opts, args := ParseFlags()

	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <github-username>")
		os.Exit(1)
	}
	username := os.Args[1]
	fmt.Printf("Fetching GitHub activity for user: %s\n", username)
	fmt.Printf("Options: %+v\n", opts)
	fmt.Println(args)
}