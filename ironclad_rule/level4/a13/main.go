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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// しゃくとり法
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// スタート地点を決める
		if i == 1 {
			r[i] = 1
		} else {
			r[i] = r[i-1]
		}

		for r[i] < n && a[r[i]+1]-a[i] <= k {
			r[i] += 1
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += (r[i] - i)
	}
	fmt.Fprintln(out, ans)
}
