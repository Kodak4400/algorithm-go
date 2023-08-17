package main

import (
	"fmt"
)

func fibo(N int) int {
	fmt.Printf("fibo(%d)を呼び出しました。\n", N)
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}

	result := fibo(N-1) + fibo(N-2)
	fmt.Printf("%d 項目 = %d\n", N, result)

	return result
}

// var memo = [6]int{0, 1, -1, -1, -1, -1}

func mfibo(n int, memo []int) int {
	fmt.Printf("func mfibo(%d) を呼びました\n", n)
	if memo[n] != -1 {
		fmt.Println("メモ化した値を返しました")
		return memo[n]
	}

	if n == 0 {
		memo[0] = 0
		return 0
	} else if n == 1 {
		memo[1] = 1
		return 1
	}

	memo[n] = mfibo(n-1, memo) + mfibo(n-2, memo)

	return memo[n]
}

func main() {
	// memo初期化
	memo := make([]int, 6)
	for i, _ := range memo {
		memo[i] = -1
	}

	mfibo(5, memo)
	fmt.Println("memo => ", memo)
}
