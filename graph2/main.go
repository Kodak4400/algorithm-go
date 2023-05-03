package main

import "fmt"

func main() {
	INF := 999999

	// ダイクストラ法
	// 2種類ある
	// - 単純に実装した場合の計算量O(|V|^2)の方法 ・・・ 蜜グラフに有利
	// - ヒープを用いる場合の計算量O(|E|log|V|)の方法 ・・・ 疎グラフに有利
	fmt.Println("== ダイクストラ法（単純） Start ==")
	m := [][][]int{
		{{1, 3}, {2, 5}},
		{{3, 12}, {2, 4}},
		{{3, 9}, {4, 4}},
		{{5, 2}},
		{{3, 7}, {5, 8}},
		{},
	}

	used := make([]bool, len(m))
	dm_dist := make([]int, len(m))
	for i, _ := range dm_dist {
		dm_dist[i] = INF
	}

	dm_dist[0] = 0
	for iter := 0; iter < len(m); iter++ {
		// 使用済み でない頂点のうち,dist値が最小の頂点を探す
		min_dist := INF
		min_v := -1
		for v := 0; v < len(m); v++ {
			if !used[v] && dm_dist[v] < min_dist {
				min_dist = dm_dist[v]
				min_v = v
			}
		}

		fmt.Println("min_v, min_dist", min_v, min_dist)

		if min_v == -1 {
			break
		}

		for _, e := range m[min_v] {
			chmin(&dm_dist[e[0]], dm_dist[min_v]+e[1])
		}
		used[min_v] = true
	}

	fmt.Println("used =>", used)
	fmt.Println("dm_dist =>", dm_dist)

	fmt.Println("== End ==")
	fmt.Println("== ダイクストラ法（ヒープ） Start ==")
	h := [][][]int{
		{{1, 3}, {2, 5}},
		{{3, 12}, {2, 4}},
		{{3, 9}, {4, 4}},
		{{5, 2}},
		{{3, 7}, {5, 8}},
		{},
	}
	dh_dist := make([]int, len(m))
	for i, _ := range dh_dist {
		dh_dist[i] = INF
	}

	// heep := make([]int, 0, len(h))

	fmt.Println("h =>", h)
	fmt.Println("dh_dist =>", dh_dist)

	fmt.Println("== End ==")
	// ベルマン・フォート法
	// 単一始点最短経路の解決策。重み付き有向グラフにおける単一始点の最短径路を解くアルゴリズムのひとつ
	// もし、始点sから到達できる負閉路が存在するなら、その旨を報告して、負閉路が存在しないなら、各頂点vへの最短経路を求めるアルゴリズム。
	// なお、すべての辺の重みが非負であれば、ダイクストラ法のほうが有効。
	fmt.Println("== ベルマン・フォート Start ==")
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
				// fmt.Println("v, t", v, t)
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
	// fmt.Println("before arr =>", arr)

	r := chmin(&arr[4], arr[1])
	fmt.Println("result =>", r)
	// fmt.Println("after arr =>", arr)
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

func push(x int, heap []int) {
	heap[len(heap)-1] = x // 最後尾に挿入
	i := len(heap) - 1    // 挿入された頂点番号取得

	for i > 0 {
		p := (i - 1) / 2 // 親の頂点番号
		if heap[p] >= x {
			break
		}
		heap[i] = heap[p]
		i = p
	}
	heap[i] = x
}

func pop(heap []int) {
	x := heap[len(heap)-1] // 頂点に持ってくる値
	i := 0                 // 根から降ろしていく
	for i*2+1 < len(heap) {
		// 子頂点同士を比較して大きい方をchild1とする
		child1 := i*2 + 1
		child2 := i*2 + 2
		if child2 < len(heap) && heap[child2] > heap[child1] {
			child1 = child2
		}
		if heap[child1] <= x {
			break
		}
		heap[i] = heap[child1]
		i = child1
	}
	heap[i] = x
}
