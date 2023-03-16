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

	fmt.Println("Flog問題 =>", dp)

	// // Flog問題（緩和）
	INF := 9999999
	var dpk = [7]int{INF, INF, INF, INF, INF, INF, INF}

	// 初期条件
	dpk[0] = 0

	for i := 1; i < 7; i++ {
		chmin(&dpk[i], dpk[i-1]+abs(h[i], h[i-1]))
		if i > 1 {
			chmin(&dpk[i], dpk[i-2]+abs(h[i], h[i-2]))
		}
	}

	fmt.Println("Flog問題（緩和） =>", dpk)

	// Flog問題（配る遷移形式）
	var dpp = [7]int{INF, INF, INF, INF, INF, INF, INF}

	// 初期条件
	dpp[0] = 0

	for i := 0; i < 7; i++ {
		if i+1 < 7 {
			chmin(&dpp[i+1], dpp[i]+abs(h[i], h[i+1]))
		}
		if i+2 < 7 {
			chmin(&dpp[i+2], dpp[i]+abs(h[i], h[i+2]))
		}
	}

	fmt.Println("Flog問題（配る遷移形式） =>", dpp)

	// dpk[0] = 8
	// dpk[1] = 5
	// chmin(&dpk[0], &dpk[1])
	// fmt.Println("Flog問題（緩和） =>", dpk)
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
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
