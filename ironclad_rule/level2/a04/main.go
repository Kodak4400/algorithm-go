package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	ans := ""

	ans = ""
	for i := 9; i >= 0; i-- {
		wari := 1 << i
		ans += strconv.Itoa((n / wari) % 2)
	}

	fmt.Fprint(out, ans)
}
