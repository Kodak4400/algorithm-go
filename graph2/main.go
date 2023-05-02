package main

import "fmt"

func main() {
	// ベルマン・フォート法
	// 単一始点最短経路の解決策。重み付き有向グラフにおける単一始点の最短径路を解くアルゴリズムのひとつ
	// もし、始点sから到達できる負閉路が存在するなら、その旨を報告して、負閉路が存在しないなら、各頂点vへの最短経路を求めるアルゴリズム。
	// なお、すべての辺の重みが非負であれば、ダイクストラ法のほうが有効。
	fmt.Println("== Start ==")
	INF := 999999
	// 頂点数:5, 辺数:12, 始点:0
	// g := [][]int{
	// 	{0, 1, 3},
	// 	{0, 3, 100},
	// 	{1, 2, 50},
	// 	{1, 3, 57},
	// 	{1, 4, -4},
	// 	{2, 4, -5},
	// 	{2, 5, 100},
	// 	{2, 3, -10},
	// 	{3, 1, -5},
	// 	{4, 3, 25},
	// 	{4, 2, 57},
	// 	{4, 5, 8},
	// }

	g := [][][]int{
		{{1, 3}, {3, 100}},
		{{2, 50}, {3, 57}, {4, -4}},
		{{4, -5}, {5, 100}, {3, -10}},
		{{1, -5}},
		{{3, 25}, {2, 57}, {5, 8}},
		{},
	}

	exist_negative_cycle := false // 負閉路を持つか
	dist := make([]int, len(g))
	for i, _ := range dist {
		dist[i] = INF
	}

	// 最初
	dist[0] = 0

	for iter := 0; iter < len(g); iter++ {
		update := false // 更新が発生したかどうかのフラグ
		for v := 0; v < len(g); v++ {
			if dist[v] == INF {
				continue
			}

			for _, t := range g[v] {
				// 暖和処理
				fmt.Println("v, t", v, t)
				if chmin(&dist[t[0]], dist[v]+t[1]) {
					update = true
				}
			}

			if !update {
				break
			}

			if iter == len(g) && update {
				exist_negative_cycle = true
			}
		}
	}

	fmt.Println(exist_negative_cycle, dist)
	fmt.Println("== End ==")

	arr := make([]int, 10)
	for i, _ := range arr {
		arr[i] = 1
	}
	arr[4] = 5
	fmt.Println("before arr =>", arr)

	r := chmin(&arr[4], arr[1])
	fmt.Println("result =>", r)
	fmt.Println("after arr =>", arr)
	// testFunc(arr)
	// fmt.Println("after arr =>", arr)
}

// 暖和を実施する。つまり制約に違反している値を取り除く
func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

// func testFunc(a []int) {
// 	a = append(a, 3)
// 	fmt.Println("add arr =>", a)
// }
