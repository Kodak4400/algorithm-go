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
	// 品物 6, 重さ 15
	//{(2,3),(1,2),(3,6),(2,1),(1,3),(5,58)}
	weight := [6]int{2, 1, 3, 2, 1, 5}
	value := [6]int{3, 2, 6, 1, 3, 58}

	np := [15]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	npz := [6][15]int{np, np, np, np, np, np}

	for i := 0; i < 5; i++ {
		for w := 0; w <= 14; w++ {
			// i 番目の品物を選ぶ場合
			if w-weight[i] >= 0 {
				chmax(&npz[i+1][w], npz[i][w-weight[i]]+value[i])
			}
			// i 番目の品物を選ばない場合
			chmax(&npz[i+1][w], npz[i][w])
		}
	}
	fmt.Println("ナップザック問題（過程） =>", npz)
	fmt.Println("ナップザック問題 =>", npz[5][14])

	// 編集距離
	s := "logistic"
	t := "algorithm"

	fmt.Println(len(s))

	// 初期化
	dph := make([][]int, len(s)+1)
	for i := range dph {
		dph[i] = make([]int, len(t)+1)
		for j := range dph[i] {
			dph[i][j] = INF
		}
	}

	dph[0][0] = 0

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			//変更操作
			if i > 0 && j > 0 {
				if string(s[i-1]) == string(t[j-1]) {
					chmin(&dph[i][j], dph[i-1][j-1])
				} else {
					chmin(&dph[i][j], dph[i-1][j-1]+1)
				}
			}
			// 削除操作
			if i > 0 {
				chmin(&dph[i][j], dph[i-1][j]+1)
			}
			// 挿入操作
			if j > 0 {
				chmin(&dph[i][j], dph[i][j-1]+1)
			}
		}
	}

	fmt.Println("編集距離問題（過程） =>", dph)
	fmt.Println("編集距離問題 =>", dph[len(s)][len(t)])

	// 区間分割の仕方を最適化
	// N個の要素が1列に並んでいて、これをいくつかの区間に分割したいものとする。
	// 各区間[l, r)にはスコアclrが付いているとする
	// KをN以下の整数として、K+1個の整数t0,t1,t2...tkを0=t0<t1<t1<....tk=Nを満たすようにとったとき、
	// 区間割[t0,t1), [t1,t2),....[tk-1, tk)のスコアをct0,t1+ct1,t2+ctK-1,tKによって定義する
	// N要素の区間分割の仕方をすべて考えたときの、考えられるスコアの最小値を求める。
	// 例
	//  N = 10
	//  K = 4
	//  t=(0, 3, 7, 8, 10)
	// この場合、スコアは、c0,3+c3,7+c7,8+c8,10となる
	//
	// 最後に区切った場所がどこであったかで、場合分けする
	// つまり、区間[0,i)の分割は、「区間[0,j)の分割に対して新たに区間[j,i)を追加したもの」とできる
	// 暖和式はchmin(dp[i], dp[j]+c[j][i])となる

	dpv := [11]int{INF, INF, INF, INF, INF, INF, INF, INF, INF, INF, INF}

	dpv[0] = 0

	c := [11][11]int{}
	c[0][3] = 1
	c[3][7] = 1
	c[7][8] = 1
	c[8][10] = 1

	for i := 0; i <= 10; i++ {
		for j := 0; j < i; j++ {
			chmin(&dpv[i], dpv[j]+c[j][i])
		}
	}

	fmt.Println("区間分割の仕方を最適化（過程） =>", dpv)
	fmt.Println("区間分割の仕方を最適化 =>", dpv[9])
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
