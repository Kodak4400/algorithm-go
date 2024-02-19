package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &q)

	var s string
	fmt.Fscan(in, &s)

	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &l[i])
		fmt.Fscan(in, &r[i])
	}

	a := make([]int, n-1)
	sum := make([]int, n)

	for i := 0; i < n-1; i++ {
		if s[i:i+1] == s[i+1:i+2] {
			a[i] = 1
		}
	}

	// 累積和
	for i := 1; i <= n-1; i++ {
		sum[i] = sum[i-1] + a[i-1]
	}

	for i := 0; i < q; i++ {
		fmt.Fprintf(out, "%d\n", sum[r[i]-1]-sum[l[i]-1])
	}
}
