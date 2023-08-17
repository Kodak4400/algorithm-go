package main

import (
	"fmt"
	"math"
)

func main() {
	var N, D int
	_, n_err := fmt.Scan(&N, &D)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	var X = make([][]int, N)
	for i := 0; i < N; i++ {
		X[i] = make([]int, D)
	}

	for j := 0; j < N; j++ {
		for i := 0; i < D; i++ {
			_, x_err := fmt.Scan(&X[j][i])
			if x_err != nil {
				fmt.Println("[Error] input.")
				return
			}
		}
	}

	fmt.Println(X)

	ans := 0
	for n := 0; n < N; n++ {
		for k := n + 1; k < N; k++ {
			sq := 0
			for j := 0; j < D; j++ {
				s := X[n][j] - X[k][j]
				sq += s * s
			}
			fmt.Println("sq =>", float64(sq))
			ss := math.Sqrt(float64(sq))
			fmt.Println("平方数 =>", ss*ss)
			if ss*ss == float64(sq) {
				ans++
			}
		}
	}

	fmt.Println("ans", ans)
}
