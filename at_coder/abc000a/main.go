package main

import (
	"bufio"
	"fmt"
	"os"
)

type St struct {
	s int
	t int
}

func searchMax(s []int) (int, int) {
	max := 0
	idx := 0
	for i, v := range s {
		if max < v {
			max = v
			idx = i
		}
	}
	return max, idx
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	st := make([]St, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &st[i].s)
		fmt.Fscan(in, &st[i].t)
	}

	for i := 0; i < n-1; i++ {
		a[i+1] += (a[i] / st[i].s) * st[i].t
	}

	fmt.Fprintln(out, a[n-1])
}
