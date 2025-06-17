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

	events := getEventsWithCache(username, opts)

	if len(events) == 0 {
		fmt.Println("No public events found for this user.")
		return
	}

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

		shown := 0
		for _, event := range events {
			if shown >= opts.Limit {
				break
			}
			if len(filterSet) > 0 && !filterSet[event.Type] {
				continue
			}
			formatted := FormatEvent(event)
			if opts.Verbose || !isUnsupportedEvent(formatted) {
				fmt.Println(formatted)
				shown++
			}
		}
		if shown == 0 {
			fmt.Println("No events match.")
		}
	}
}

func getEventsWithCache(username string, opts *CLIOptions) []GitHubEvent {
	cache, err := LoadCache()
	if err != nil {
		fmt.Printf("Warning: could not load cache: %v\n", err)
		cache = Cache{}
	}

	if !opts.Refresh {
		if cached, ok := GetCachedEvents(cache, username); ok {
			fmt.Println("(Loaded from cache)")
			return cached
		}
	}

	events, err := FetchUserEvents(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	SetCachedEvents(cache, username, events)
	if err := SaveCache(cache); err != nil {
		fmt.Printf("Warning: could not save cache: %v\n", err)
	}
	return events
}
