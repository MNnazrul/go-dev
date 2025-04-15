package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	ans, sm, mn := -1, 0, gas[0] - cost[0]
	n := len(gas)

	for idx := 0; idx < n; idx++ {
		sm += gas[idx] - cost[idx]
		if mn < sm {
			mn = sm
			ans = idx
		}
	}

	if sm < 0 {
		return -1
	}

	sm = 0
	ans1 := ans
	for i := ans; i >= 0; i-- {
		if gas[i] - cost[i] < 0 {
			return i + 1
		}
		ans1 = i
	}

	for i := len(cost) - 1; i > ans; i-- {
		if gas[i] - cost[i] < 0 {
			if i == len(cost) - 1 {
				return 0
			} else {
				return i + 1
			}
		}
		ans1 = ans
	}

	return ans1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))
}