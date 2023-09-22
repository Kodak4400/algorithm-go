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

func max(max int, i int) int {
	if max < i {
		max = i
	}
	return max
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
	for i0 := 0; i0 < m; i0++ {
		for i1 := 0; i1 < m; i1++ {
			for i2 := 0; i2 < m; i2++ {
				if a1[i0] != a2[i1] {
					continue
				}
				if a1[i0] != a3[i2] {
					continue
				}
				now := 0
				now = i0
				{
					t := i1
					if i0 == i1 {
						t += m
					}
					now = max(now, t)
				}
				{
					t := i2
					if i0 == i2 {
						t += m
					}
					if i1 == i2 {
						t += m
					}
					now = max(now, t)
				}
				ans = min(ans, now)
			}
		}
	}
	if ans == INF {
		ans = -1
	}
	fmt.Fprintln(out, ans)
}
