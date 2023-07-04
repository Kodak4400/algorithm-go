package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var as []uint64
	for i := 0; i < 64; i++ {
		var a uint64
		fmt.Fscan(r, &a)
		as = append(as, a)
	}

	var ans uint64
	for i, v := range as {
		ans += v * uint64(math.Pow(2, float64(i)))
	}
	fmt.Println(ans)
}
