package main

import (
	"fmt"
	"os"
)

func main() {
	
	opts, args := ParseFlags()

	if len(args) < 1 {
		fmt.Println("Usage: github-activity <github-username>")
		os.Exit(1)
	}
	username := args[0]
	fmt.Printf("Fetching GitHub activity for user: %s\n", username)
	fmt.Printf("options : %+v\n", opts)
	
	events, err := FetchUserEvents(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Fetched %d events.\n", len(events))

	if opts.Raw {
		for _, event := range events {
			fmt.Printf("%+v\n\n", event)
		}
	}

}