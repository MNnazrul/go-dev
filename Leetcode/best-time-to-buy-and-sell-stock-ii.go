package main

import "fmt"

func maxProfite(prices []int) int {
	ans := 0
	for i:=1 ; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			ans += prices[i] - prices[i - 1]
		}
	}

	return ans
}

func main() {
	fmt.Println(maxProfite([]int{7,1,5,3,6,4}))
}

/*



*/