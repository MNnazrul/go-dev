func search(nums []int, target int) int {
    
    
    chk := func ( l, r int) int {
        for l <= r {
            md := (l + r) >> 1
            if nums[md] == target {
                return md
            }
            if nums[md] < target {
                l = md + 1
            } else {
                r = md - 1
            }
        }
        return -1
    }

    l, r, ans := 0, len(nums) - 1, len(nums)
    for l <= r {
        md := (l + r) / 2
        if nums[md] >= nums[0] {
            l = md + 1
        } else {
            r = md - 1 
            ans = min(ans, md)
        }
    }
    return  max(chk(0, ans%len(nums) - 1), chk(ans%len(nums), len(nums) - 1))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}