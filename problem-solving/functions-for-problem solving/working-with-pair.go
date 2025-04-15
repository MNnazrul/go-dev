package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	first  int
	second int
}

func main() {
	pairs := []Pair{
		{1, 5},
		{2, 3},
		{1, 9},
		{3, 7},
		{2, 4},
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].first == pairs[j].first {
			
			return pairs[i].second > pairs[j].second
		}
		
		return pairs[i].first < pairs[j].first
	})

	for _, p := range pairs {
		fmt.Printf("(%d, %d) ", p.first, p.second)
	}
}
