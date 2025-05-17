func maxProduct(nums []int) int {
    ans, mul, mn := -1000000000, 1, -1000000000
    for _, val := range nums {
        mul *= val
        if mul == 0 {
            mul = 1
            mn = -1000000000
            ans = max(ans, 0)
            continue
        }
        if mul < 0 && mn >= mul {
            ans = max(ans, mul / mn)
        }
        ans = max(ans, mul)
        if mul < 0 && mn < mul {
            mn = mul
        }
    }
    return ans
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}