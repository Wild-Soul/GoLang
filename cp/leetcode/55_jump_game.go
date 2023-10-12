func canJump(nums []int) bool {
    n := len(nums);
    vis := make([]int, n);

    vis[0] = 1;
    if n > 1 {
        vis[1] = -1;
    }

    // if we can reach position i, then we can jump to position i+1.
    // by default we've already reach position 0.

    for i:=0; i<n; i++ {
        
        if i>0 {
            vis[i] += vis[i-1];
        }
        
        if vis[i] > 0 {
            if i+1 < n {
                vis[i+1]++;
            }
            if i+nums[i]+1 < n {
                vis[i+nums[i]+1]--;
            }
        }
        
        // for i:=0; i<n; i++ {
        //     fmt.Print(vis[i], " ");
        // }
        // fmt.Println();
    }
    
    return vis[n-1] > 0;
}

func min(a, b int) int {
    if a < b {
        return a;
    }
    return b;
}

