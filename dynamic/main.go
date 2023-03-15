package main

import (
	"fmt"
)

func main() {
	//Frog問題
	// N = 7
	// h = {2, 9, 4, 5, 1, 6, 10}
	// cst1 = {7, 5, 1, 4, 5, 4}
	// cst2 = {2, 4, 3, 1, 9}

	var dp = [7]int{0, 0, 0, 0, 0, 0, 0}
	h := [7]int{2, 9, 4, 5, 1, 6, 10}

	for i := 1; i < 7; i++ {
		if i == 1 {
			dp[i] = abs(h[i], h[i-1])
		} else {
			dp[i] = min(dp[i-1]+abs(h[i], h[i-1]), dp[i-2]+abs(h[i], h[i-2]))
		}
	}

	fmt.Println("答え =>", dp)
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
