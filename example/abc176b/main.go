package main

import (
	"fmt"
)

func main() {
	var N, K int
	_, n_err := fmt.Scan(&N, &K)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	ans := 0
	for 0 < N {
		N /= K
		ans++
	}
	fmt.Println(ans)
}
