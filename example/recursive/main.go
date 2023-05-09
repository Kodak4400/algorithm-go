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

func main() {
	fibo(5)
}
