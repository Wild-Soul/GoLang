func climbStairs(n int) int {
    var a,b,c int;
    a = 1; // number of ways to reach step 1
    b = 2; // number of ways to reach step 2
    if n==1 {
        return a
    } else if n==2 {
        return b;
    }
    for i:=3; i<=n; i++ {
        c = a+b; // number of ways to reach step n = number of ways to reach step n-2 + n-1
        a = b;
        b = c;
    }
    return b;
}

