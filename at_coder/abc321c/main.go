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

	var k int
	fmt.Fscan(in, &k)

	var a = make([]int, 0, 1<<10)
	for i := 0; i < 1<<10; i++ {
		var x = 0
		for j := 9; j >= 0; j-- {
			if i>>j&1 == 1 {
				x = x*10 + j
			}
		}
		if x == 0 {
			continue
		}
		a = append(a, x)
	}
	sort.Ints(a)

	fmt.Fprintln(out, a[k-1])
}
