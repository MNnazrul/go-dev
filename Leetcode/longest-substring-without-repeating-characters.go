package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    
    var (
		ans int = 0
		mx int = 0
	)

	charMap := make(map[rune]int)

    for index, value := range s {
		val, _ := charMap[value]
		
		if mx < val  {
			mx = val 
		}
        if ans < index + 1 - mx {
			ans = index + 1 - mx
        }
		charMap[value] = index + 1
    }
    
    return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pww kew 3"))
}