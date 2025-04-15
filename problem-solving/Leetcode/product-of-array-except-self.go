package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	prev := 1

	for idx, value := range(nums) {
		answer[idx] = prev
		prev *= value
	}

	prev = 1
	for idx := n - 1; idx >= 0; idx-- {
		answer[idx] *= prev
		prev *= nums[idx]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}