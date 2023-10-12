func minCostClimbingStairs(cost []int) int {
    n := len(cost);
    
    if n==1 {
        return cost[0];
    } else if n==2 {
        return int(math.Min( float64(cost[0]), float64(cost[1])));
    }
    
    dp := make([]int, n+1);
    MAX_VAL := 1e9 + 1;
    for i:=0; i<=n; i++ {
        dp[i] = int(MAX_VAL);
    }
    
    dp[0] = 0; dp[1] = 0;
    for i:=0; i<n; i++ {
        // if we pay cost here then we can reach i+1 and i+2 with this cost.
        if i+1 <=n {
            dp[i+1] = int(math.Min(float64(dp[i+1]), float64(dp[i] + cost[i])));
        }
        if i+2 <= n {
            dp[i+2] = int(math.Min( float64(dp[i+2]), float64(dp[i] + cost[i])));
        }
    }
    return dp[n];
}

