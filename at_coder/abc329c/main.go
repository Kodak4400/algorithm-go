package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)

	ss := strings.Split(s, "")
	m := make(map[string]int)

	left := 1
	m[ss[0]] = 1
	ex := false
	for i := 0; i < n-1; i++ {
		if ss[i] == ss[i+1] {
			left++
		} else {
			ex = true
		}

		// fmt.Println("t", ss[i], ss[i+1], left, ss[i] == ss[i+1])

		if m[ss[i]] < left {
			m[ss[i]] = left
		}

		if ex {
			left = 1
			ex = false
		}
	}

	ans := 0
	for _, v := range m {
		ans += v
	}

	fmt.Fprintln(out, ans)
}
