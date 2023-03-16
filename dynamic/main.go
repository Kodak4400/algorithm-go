package main

import (
	"fmt"
)

var INF = 9999999
var h = [7]int{2, 9, 4, 5, 1, 6, 10}

// Flog問題（再帰）
var dpr = [7]int{INF, INF, INF, INF, INF, INF, INF}

func main() {
	//Frog問題
	// N = 7
	// h = {2, 9, 4, 5, 1, 6, 10}
	// cst1 = {7, 5, 1, 4, 5, 4}
	// cst2 = {2, 4, 3, 1, 9}

	var dp = [7]int{0, 0, 0, 0, 0, 0, 0}
	// h := [7]int{2, 9, 4, 5, 1, 6, 10}

	for i := 1; i < 7; i++ {
		if i == 1 {
			dp[i] = abs(h[i], h[i-1])
		} else {
			dp[i] = min(dp[i-1]+abs(h[i], h[i-1]), dp[i-2]+abs(h[i], h[i-2]))
		}
	}

	fmt.Println("Flog問題 =>", dp)

	// // Flog問題（緩和）
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

	// Flog問題（再帰）
	rec(6)
	fmt.Println("Flog問題（再帰） =>", dpr)

	// ナップザック問題
	// 品物 6 (weight, value), 重さ we
	//{(2,3),(1,2),(3,6),(2,1),(1,3),(5,58)}
	npz := [6][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	// weight := [6]int{2, 1, 3, 2, 1, 5}
	// value := [6]int{3, 2, 6, 1, 3, 58}

	for i := 0; i < 5; i++ {
		for w := 0; w <= 1; w++ {
			// i 番目の品物を選ぶ場合
			// if w-weight[i] >= 0 {
			// 	chmax(&npz[i+1][w], npz[i][w-weight[i]]+value[i])
			// }
			// i 番目の品物を選ばない場合
			chmax(&npz[i+1][w], npz[i][w])
		}
	}
	fmt.Println("ナップザック問題 =>", npz)
}

func rec(i int) int {
	if dpr[i] < INF {
		return dpr[i]
	}

	if i == 0 {
		dpr[i] = 0
		return 0
	}

	res := INF

	// 足場 i - 1 から来た場合
	chmin(&res, rec(i-1)+abs(h[i], h[i-1]))

	// 足場 i - 2 から来た場合
	if i > 1 {
		chmin(&res, rec(i-2)+abs(h[i], h[i-2]))
	}

	dpr[i] = res
	return dpr[i]
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
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
