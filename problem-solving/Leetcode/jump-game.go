package main

import "fmt"

func canJump(nums []int) bool {
	
	sz := len(nums)
	mx := sz - 1
	for i := sz - 1; i >= 0; i-- {
		if i + nums[i] >= mx {
			mx = i
		}
	}

	if mx == 0 {
		return true
	}

	return false
}


func main() {
	fmt.Println(canJump([]int{0}))
}

