package main

import (
	"fmt"
	"os"

	"github.com/MNnazrul/caching-proxy/cache"
	"github.com/MNnazrul/caching-proxy/cli"
	"github.com/MNnazrul/caching-proxy/proxy"
)

func main() {
	args := cli.ParseArgs()
	memCache := cache.NewMemoryCache()

	if args.ClearCache {
		memCache.Clear()
		fmt.Println("Cache clearer (Placeholder).")
		os.Exit(0)
	}
	fmt.Printf("Starting proxy on port %d, origin: %s\n", args.Port, args.Origin)
	err := proxy.StartProxy(args.Port, args.Origin, memCache)
	if err != nil {
		fmt.Println("Error starting proxy:", err)
		os.Exit(1)
	}
}
