package main 

import (
	."fmt"
)

func main() {
	Println(findMin([]int{4,5,6,7,0,1,2}))
}

func findMin(nums []int) int {
    l, r, ans := 0, len(nums) - 1, len(nums)
    for l <= r {
        md := (l + r) / 2
        Println("mid", md, "l", l, "r", r)
        if nums[md] >= nums[0] {
            l = md + 1
        } else {
        	Println(l, r, md)
            r = md - 1 
            ans = min(ans, md)
        }
    }
    return nums[ans]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}