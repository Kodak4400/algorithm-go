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

	var d, n int
	fmt.Fscan(in, &d)
	fmt.Fscan(in, &n)

	l := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
	}

	ans := make([]int, d+1)
	for i := 1; i <= n; i++ {
		for j := l[i]; j <= r[i]; j++ {
			ans[j]++
		}
	}

	for i := 1; i <= d; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
