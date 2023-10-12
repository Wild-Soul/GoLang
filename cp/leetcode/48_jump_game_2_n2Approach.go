func jump(nums []int) int {
    n := len(nums);
    
    // a helper function, contains business logic.
    // idea -> from each position i, update the positions that can be reached from this
    // while updating the positon make sure to use minimum of existing value and the update value.
    // time complexity -> O(n*n); can be optimized.
    solve := func (n int) int {
        dp := make([]int, n);
        for i:=0; i<n; i++ {
            dp[i] = math.MaxInt32
        }
        dp[0] = 0;
        for i:=0; i<n; i++ {
            for j:=1; j<=nums[i] && (i+j) < n; j++ {
                dp[i+j] = int(math.Min(float64(dp[i+j]), float64(dp[i]+1)));
            }
        }
        return dp[n-1];
    }
    return solve(n);
}

