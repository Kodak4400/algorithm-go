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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &l[i])
		fmt.Fscan(in, &r[i])
	}

	// fmt.Println(a, l, r)

	ruiseki := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ruiseki[i] = ruiseki[i-1] + a[i]
	}

	// fmt.Println(ruiseki)
	for i := 0; i < q; i++ {
		ans := ruiseki[r[i]] - ruiseki[l[i]-1]
		fmt.Fprintln(out, ans)
	}
}
