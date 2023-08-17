package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(mid int, k int, n int, a []int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += mid / a[i]
	}

	if sum >= k {
		return true
	}

	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	left := 1
	right := int(1e9)

	for left < right {
		mid := (left + right) / 2
		ans := check(mid, k, n, a)
		if ans {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Fprintln(out, left)
}
