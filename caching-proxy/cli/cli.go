package cli

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Port       int
	Origin     string
	ClearCache bool
}

func ParseArgs() Args {
	port := flag.Int("port", 0, "Port to run the proxy server on (required)")
	origin := flag.String("origin", "", "Origin server URL (required)")
	clearCache := flag.Bool("clear-cache", false, "Clear the cache and exit")

	flag.Parse()

	if *port == 0 || *origin == "" {
		fmt.Println("Error: --port and --origin are required")
		flag.Usage()
		os.Exit(1)
	}

	return Args{
		Port:       *port,
		Origin:     *origin,
		ClearCache: *clearCache,
	}
}
