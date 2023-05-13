package main

import (
	"fmt"
)

func main() {
	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	// i * j = N

	res := false

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == N {
				res = true
				break
			}
		}
	}

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
