package main

import (
	"fmt"
)

func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if sum == 1 {
		return 10
	}
	return sum
}

func main() {
	// 2 <= A <= B <= 100000
	// A + B = N
	// 各位の和は A + N > Nとなる。
	// 15 => 10 + 5
	// 100000 => 10000 + 90000
	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	fmt.Println(findSumOfDigits(N))
}
