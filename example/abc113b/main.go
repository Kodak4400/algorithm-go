package main

import (
	"fmt"
	"math"
)

func Abs(i, j int) int {
	if i < j {
		return j - i
	}
	return i - j
}

func main() {
	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	var T, A int
	_, ta_err := fmt.Scan(&T, &A)
	if ta_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	H := make([]int, N, N)
	for i := 0; i < N; i++ {
		_, h_err := fmt.Scan(&H[i])
		if h_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sav := math.MaxInt
	res := 1
	for i := 0; i < N; i++ {
		sum := int(T - H[i]*6/1000)
		if sav > Abs(T, sum) {
			sav = Abs(T, sum)
			res = i + 1
		}
	}
	fmt.Println(res)
}
