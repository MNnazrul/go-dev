package main

import (
	"fmt"
	"os"

	"github.com/MNnazrul/caching-proxy/cli"
)

func main() {
	args := cli.ParseArgs()

	if args.ClearCache {
		// Implement cache clearer logic
		fmt.Println("Cache clearer (Placeholder).")
		os.Exit(0)
	}
	fmt.Printf("Starting proxy on port %d, origin: %s\n", args.Port, args.Origin)
}
