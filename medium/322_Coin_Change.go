func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    dp[0] = 0
    for i := 1; i < amount+1; i++ {
        dp[i] = amount+1
    }

    for i := 0; i < amount+1; i++ {
        for _, c := range coins {
            if i-c < 0 {
                continue
            }
            dp[i] = min(dp[i], dp[i-c]+1)
        }
    }

    if dp[amount] > amount {
        return -1
    }

    return dp[amount]
}

func min (a, b int) int {
    if a >= b {
        return b
    }
    return a
}


