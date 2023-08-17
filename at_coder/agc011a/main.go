package main

import (
	"fmt"
	"sort"
)

func main() {
	// N人 飛行機到着
	// C バスの定員
	// K バス待ち時間
	// T[i] 時刻に到着
	// T[i]以上T[i]+K以下のバスに乗れないとダメ
	// バスの最小値を求める
	// バスの出発時刻は整数でなくても良い、同じ時刻に出発するバスが複数あっても良い

	// 7 3 4
	// 1
	// 2 2 - 1 = 1
	// 3 3 - 1 = 2 (制限 3)
	// 4
	// 6 6 - 4 = 2
	// 9 8 - 4 = 4 (制限 3, 時間 4)
	// 12

	var N, C, K int
	_, nck_err := fmt.Scan(&N, &C, &K)
	if nck_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	T := make([]int, N)
	for i := 0; i < N; i++ {
		_, a_err := fmt.Scan(&T[i])
		if a_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(T, func(i, j int) bool { return T[i] < T[j] })

	G := make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = make([]int, C)
	}
	// [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]

	// g := 0
	// for i := 0; i < N; i++ {
	// 	G[g][0]
	// }

}
