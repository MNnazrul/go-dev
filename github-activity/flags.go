package main

import (
	"flag"
)

type CLIOptions struct {
	Filter  string
	Limit   int
	Refresh bool
	Raw     bool
}

func ParseFlags() (*CLIOptions, []string) {
	opts := &CLIOptions{}

	flag.StringVar(&opts.Filter, "filter", "", "Filter by event types (comma-separated)")
	flag.StringVar(&opts.Filter, "f", "", "Filter by event types (shorthand)")
	flag.IntVar(&opts.Limit, "limit", 10, "Number of events to display (max 100)")
	flag.IntVar(&opts.Limit, "l", 10, "Number of events to display (shorthand)")
	flag.BoolVar(&opts.Refresh, "refresh", false, "Skip cache and force API request")
	flag.BoolVar(&opts.Raw, "raw", false, "Show raw JSON output")

	flag.Parse()
	return opts, flag.Args()
}
