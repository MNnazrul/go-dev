func missingNumber(nums []int) int {
    sum := len(nums) * (len(nums) + 1) / 2
    for _, val := range nums {
        sum -= val
    }
    return sum
}

func init() {
    debug.SetMemoryLimit(1)
}