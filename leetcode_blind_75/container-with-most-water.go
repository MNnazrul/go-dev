func maxArea(height []int) int {
    l, r, ans := 0, len(height) - 1, 0
    for l <= r {
        ans = max(ans, min(height[l], height[r]) * (r - l))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}