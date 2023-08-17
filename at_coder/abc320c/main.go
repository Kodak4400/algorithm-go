package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var m int
	fmt.Fscan(in, &m)

	var s1, s2, s3 string
	fmt.Fscan(in, &s1)
	fmt.Fscan(in, &s2)
	fmt.Fscan(in, &s3)

	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")
	a3 := strings.Split(s3, "")

	INF := 1001001001
	ans := INF
	for t0 := 0; t0 < 300; t0++ {
		for t1 := 0; t1 < 300; t1++ {
			for t2 := 0; t2 < 300; t2++ {
				if t0 == t1 {
					continue
				}
				if t0 == t2 {
					continue
				}
				if t1 == t2 {
					continue
				}
				if a1[t0%m] != a2[t1%m] {
					continue
				}
				if a1[t0%m] != a3[t2%m] {
					continue
				}
				ans = min(ans, tMax(t0, t1, t2))
			}
		}
	}
	if ans == INF {
		ans = -1
	}
	fmt.Fprintln(out, ans)
}
