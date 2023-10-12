func tribonacci(n int) int {
    if n<2 {
        return n;
    }
    if n==2 {
        return 1;
    }
    a:=0;
    b:=1;
    c:=1;
    for i:=3; i<=n; i++ {
        d:= a+b+c;
        a = b;
        b = c;
        c = d;
    }
    return c;
}
