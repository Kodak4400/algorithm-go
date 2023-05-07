package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("N枚のカードがあり、i枚目のカードには aiという数が書かれています。\nAlice と Bob はこれらのカードを使ってゲームを行います。ゲームでは 2 人が交互に 1 枚ずつカードを取っていきます。Alice が先にカードを取ります。\n2人がすべてのカードを取ったときゲームは終了し、取ったカードの数の合計がその人の得点になります。2 人とも自分の得点を最大化するように最適戦略をとったとき、Alice は Bob より何点多くの得点を獲得できるかを求めてください。")

	var N int
	fmt.Print("N:")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	A := make([]int, N, N)
	fmt.Print("A:")
	for i := 0; i < N; i++ {
		_, a_err := fmt.Scan(&A[i])
		if a_err != nil {
			fmt.Println("err: INPUT.")
			return
		}
	}

	// 安定じゃないソート
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	Alice := 0
	Bob := 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			Alice += A[i]
		} else {
			Bob += A[i]
		}
	}

	fmt.Println(Alice - Bob)
}
