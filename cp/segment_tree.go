package main

import (
    "fmt";
    "math"
)

func min (a,b int) int {
    if a<b {
        return a;
    }
    return b;
}

func buildSt (curNode, l, r, n int, st []int, a []int) {
    if l > r {
        return;
    }
    if (l == r) {
        st[curNode] = a[l];
        return;
    }
    mid := l + (r-l)/2;
    buildSt (2*curNode, l, mid, n, st, a);
    buildSt (2*curNode+1, mid+1, r, n, st, a);
    st[curNode] = min( st[2*curNode], st[2*curNode+1]);
}

func query (curNode, l, r, qL, qR, n int, st []int) int {
    if r < qL || l > qR {
        return math.MaxInt32;
    }
    // if current node fits perfectly in the query range then just return the value.
    if l>=qL && r<=qR {
        return st[curNode];
    }
    mid := l + (r-l)/2;
    return min(
        query(2*curNode, l, mid, qL, qR, n, st),
        query(2*curNode+1, mid+1, r, qL, qR, n , st),
    );
}

func updateSt (curNode, l, r, pos, val int, st []int) {
    if l>r {
        return;
    }
    if l==r {
        st[curNode] = val;
        return;
    }

    mid := l + (r-l)/2;

    if l <= pos && pos <= mid {
        updateSt (2*curNode, l, mid, pos, val, st);
    } else {
        updateSt (2*curNode+1, mid+1, r, pos, val, st);
    }
    st[curNode] = min (st[curNode*2], st[curNode*2+1]);
    return;
}

func main() {
    var n,q,x,y int;
    var ch byte;
    fmt.Scanf("%d %d",&n,&q);
    // fmt.Scan(&n);
    // fmt.Scan(&q);

    // fmt.Println(n,q);
    a := make([]int, n+1);
    st := make([]int, 4*n+4);

    for i:=1; i<=n; i++ {
        fmt.Scan(&a[i]);
    }
    // fmt.Println(a);

    // build segment tree.
    buildSt(1, 1, n, n, st, a);

    // fmt.Println(st);

    for ; q>0; q-- {
        fmt.Scanf("%c %d %d", &ch, &x, &y);
        if ch == 'q' {
            ans := query(1, 1, n, x, y, n, st);
            fmt.Println(ans);
        } else {
            updateSt(1, 1, n, x, y, st);
        }
    }
}

