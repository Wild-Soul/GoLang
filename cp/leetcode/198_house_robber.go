func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}


func rob(nums []int) int {
    n := len(nums)
    if len(nums) == 0 {
        return 0    
    }
    
    a,b := 0,0
    ans := 0
    for i:=0; i<n; i++ {
        if (a+nums[i]) > ans {
            ans = a + nums[i]
            b,a = max(a+nums[i],b),max(a,b)
        } else {
            a = max(a,b)
        }
    }
    
     return ans
}

