package main

import (
	"fmt"
)

var memo = [7]int{-1, -1, -1, -1, -1, -1, -1}

func main() {
	// フィボナッチ数列を求める
	fmt.Println("fibo(6) =>", fibo(6))

	// フィボナッチ数列を求める(for版)
	fibofor := [7]int{0, 1}
	for n := 2; n < 7; n++ {
		fibofor[n] = fibofor[n-1] + fibofor[n-2]
		fmt.Printf("%d 項目 = %d\n", n, fibofor[n])
	}

	// フィボナッチ数列を求める(メモ化)
	fmt.Println("mfibo(6) =>", mfibo(6))
}

// ------------------------
func fibo(n int) int {
	fmt.Printf("func fibo(%d) を呼びました\n", n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := fibo(n-1) + fibo(n-2)
	fmt.Printf("%d 項目 = %d\n", n, result)

	return result
}

func mfibo(n int) int {
	fmt.Printf("func mfibo(%d) を呼びました\n", n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = mfibo(n-1) + mfibo(n-2)

	return memo[n]
}
