package main

import (
	"fmt"
)

func main() {

	// 再帰と分割統治法
	p(5)

	// ユークリッドの互除法で最大公約数を求める
	fmt.Println("gcd(51, 15) =>", gcd(51, 15))
	fmt.Println("gcd(15, 51) =>", gcd(15, 51))
}

// ------------------------
func p(n int) int {
	fmt.Printf("func p(%d) を呼びました\n", n)

	if n == 0 {
		return 0
	}

	result := n + p(n-1)
	fmt.Printf("%d までの和 = %d\n", n, result)
	return result
}

func gcd(m int, n int) int {
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}
