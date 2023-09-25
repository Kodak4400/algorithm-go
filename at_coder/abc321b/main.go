package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func reverse(s string) string {
	runes := []rune(s) // rune型はcode pointを単位として扱う型
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(min int, i int) int {
	if min > i {
		min = i
	}
	return min
}

func tMax(t0 int, t1 int, t2 int) int {
	max := -1
	if t0 < t1 {
		if t1 < t2 {
			max = t2
		} else {
			max = t1
		}
	} else {
		if t0 < t2 {
			max = t2
		} else {
			max = t0
		}
	}
	return max
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &x)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := -1
	for i := 0; i <= 100; i++ {
		// ソートはコピーした配列に対して行うこと
		copy(b, a)
		b[n-1] = i
		sort.Ints(b)

		sum := 0
		for j := 1; j < n-1; j++ {
			sum += b[j]
		}

		if x <= sum {
			ans = i
			break
		}
	}

	fmt.Fprintln(out, ans)
}
