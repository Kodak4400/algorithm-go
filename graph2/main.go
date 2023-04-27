package main

import "fmt"

func main() {
	// ベルマン・フォート法
	// 単一始点最短経路の解決策。重み付き有向グラフにおける単一始点の最短径路を解くアルゴリズムのひとつ
	// もし、始点sから到達できる負閉路が存在するなら、その旨を報告して、負閉路が存在しないなら、各頂点vへの最短経路を求めるアルゴリズム。
	// なお、すべての辺の重みが非負であれば、ダイクストラ法のほうが有効。
	fmt.Println("== Start ==")
	// Golangの練習
	arr := make([]int, 0, 5)
	fmt.Println("before arr =>", arr)
	testFunc(arr)
	fmt.Println("after arr =>", arr)
}

func testFunc(a []int) {
	a = append(a, 3)
	fmt.Println("add arr =>", a)
}
