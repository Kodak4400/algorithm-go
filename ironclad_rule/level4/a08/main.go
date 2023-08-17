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

	var h, w int
	fmt.Fscan(in, &h, &w)

	x := make([][]int, h)
	for i := 0; i < w; i++ {
		x[i] = make([]int, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &x[i][j])
		}
	}

	var q int
	fmt.Fscan(in, &q)

	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i], &d[i])
	}

	fmt.Println(x, a, b, c, d)

}
