func wordBreak(s string, wordDict []string) bool {
    var dp func(id int) bool
    mem := make([]int, len(s) + 1)
    for i := 0; i <= len(s); i++ {
        mem[i] = -1
    }
    dp = func(id int) bool {
        if id == len(s) {
            return true
        }
        if mem[id] != -1 {
            if mem[id] == 1 {
                return true
            } else {
                return false
            }
        }
        flg := false
        for _, st := range wordDict {
            if strings.HasPrefix(s[id:], st) {
                if dp(id + len(st)) {
                    flg = true
                    break
                }
            }
        }
        if flg == true {
            mem[id] = 1
        } else {
            mem[id] = 0
        }
        return flg
    }    

    return dp(0)
}