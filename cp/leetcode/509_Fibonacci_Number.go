func fib(n int) int {
    a:=0;
    b:=1;
    if(n<2) {
        return n;
    }
    for i:=2; i<=n; i++ {
        c:=a+b;
        a=b;
        b=c;
    }
    return b;
}

