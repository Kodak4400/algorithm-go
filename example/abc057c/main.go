package main

import (
	"fmt"
	"strconv"
)

// func F(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func main() {
	// 10000
	// F(10000, 0) => 5
	// F(1000, 10) => 4
	// F(100, 100) => 3

	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	n := len(strconv.Itoa(N))

	// 全探索
	// ans := math.MaxInt64
	for i := 0; i < n; i++ {
		A := N
		B := i

		a := len(strconv.Itoa(A))
		b := len(strconv.Itoa(B))

		if a*b == N {
			fmt.Println(a, b)
		}

		// ans = F(len(strconv.Itoa(A)), len(strconv.Itoa(B)))
		N /= 10
	}
}
