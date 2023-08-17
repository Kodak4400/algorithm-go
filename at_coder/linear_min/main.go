package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("N個の整数a0,a1....aN-1とN個の整数b0,b1...bN-1が与えられる。\n2組の正数列からそれぞれ1個ずつ整数を選んで和を取ります。その和として考えられる値のうち、整数K以上の範囲内で最小値を求めてください。")

	var N int
	fmt.Print("N:")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("input error.")
		return
	}

	a := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Printf("a[%d]:", i)
		_, a_err := fmt.Scan(&a[i])
		if a_err != nil {
			fmt.Println("input error.")
			return
		}
	}

	b := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Printf("b[%d]:", i)
		_, b_err := fmt.Scan(&b[i])
		if b_err != nil {
			fmt.Println("input error.")
			return
		}
	}

	var K int
	fmt.Print("K:")
	_, k_err := fmt.Scan(&K)
	if k_err != nil {
		fmt.Println("input error.")
		return
	}

	pair_min_value := math.MaxInt64
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum := a[i] + b[j]
			if sum <= K {
				continue
			}
			if sum < pair_min_value {
				pair_min_value = sum
			}
		}
	}

	fmt.Println(pair_min_value)
}
