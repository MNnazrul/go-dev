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
	} else {

		filterSet := map[string]bool{}
		if opts.Filter != "" {
			for _, t := range splitAndTrim(opts.Filter) {
				filterSet[t] = true
			}
		}

		count := opts.Limit
		for _, event := range events {
			if count == 0 {
				break
			}
			if len(filterSet) > 0 && !filterSet[event.Type] {
				continue
			}
			count--
			fmt.Println(FormatEvent(event))
		}

	}
}
