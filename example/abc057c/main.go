package main

import (
	"fmt"
	"strconv"
)

func F(A, B string) int {
	a := len(A)
	b := len(B)
	if a < b {
		return b
	}
	return a
}

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

	ans := len(strconv.Itoa(N)) // 最大の桁数

	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		j := N / i
		cur := F(strconv.Itoa(i), strconv.Itoa(j))

		if ans > cur {
			ans = cur
		}

	}
	fmt.Println(ans)
}
