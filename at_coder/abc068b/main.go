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

	macnt, ans := -1, N

	for i := N; i >= 1; i-- {
		cnt := 0
		y := i
		for y%2 == 0 {
			y /= 2
			cnt++
			if macnt < cnt {
				macnt = cnt
				ans = i
			}
		}
	}
	fmt.Println(ans)
}
