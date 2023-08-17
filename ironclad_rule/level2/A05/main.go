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
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)

	cnt := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if k-(i+j) >= 1 && n >= k-(i+j) {
				cnt++
			}
		}
	}

	fmt.Fprint(out, cnt)
}
