func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}

func solve(nums []int, st, end int) int {
    if end <= st {
        return 0;
    }
    a,b := 0,0
    ans := 0
    for i:=st; i<end; i++ {
        if (a+nums[i]) > ans {
            ans = a + nums[i]
            b,a = max(a+nums[i],b),max(a,b)
        } else {
            a = max(a,b)
        }
    }
    return ans;
}

func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0    
    }
    
    ans := max ( solve(nums, 1, n), nums[0] + solve(nums, 2, n-1) );
    
    return ans
}

