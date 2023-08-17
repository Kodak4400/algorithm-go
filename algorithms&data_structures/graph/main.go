package main

import "fmt"

// var color = []int{-1, -1, -1, -1, -1}

func main() {
	//...
	// グラフの表現（有効グラフ）
	// G[頂点] = {隣接点}
	// G[0] = {5}
	// G[1] = {3,6}
	// G[2] = {5,7}
	// G[3] = {0,7}
	// G[4] = {1,2,6}
	// G[5] = {}
	// G[6] = {7}
	// G[7] = {0}

	g := [][]int{
		{5},
		{3, 6},
		{5, 7},
		{0, 7},
		{1, 2, 6},
		{},
		{7},
		{0},
	}
	// グラフGにおいて、頂点sを始点とした探索を行う
	search(g, 4)
	dsearch(g, 4)

	h := [][]int{
		{1, 4, 2},
		{0, 3, 4, 8},
		{0, 5},
		{1, 7, 8},
		{0, 1, 8},
		{2, 6, 8},
		{5, 7},
		{3, 6},
		{1, 3, 4, 5},
	}
	hsearch(h, 0)

	// 2部グラフ
	// 白色の頂点同士が隣接することなく、黒色の頂点同士が隣接することもない。
	// という条件を満たす各頂点を白色または黒色に塗り分けることが可能なグラフのこと
	// 与えられたグラフが2部グラフかどうかを判定する。
	// 以下を繰り返し、この過程で「両端点が同色であるような辺」が検出されれば2部グラフではないことが確定する。
	// - 白色頂点に隣接した頂点は黒色
	// - 黒色頂点に隣接した頂点は白色
	j := [][]int{
		{1, 3},
		{0, 2, 4},
		{1},
		{0, 4},
		{1, 3},
	}

	color := make([]int, len(j))
	for i, _ := range color {
		color[i] = -1
	}

	is_bipartite := true
	for v := 0; v < len(j); v++ {
		if color[v] != -1 {
			continue
		}
		if !dfs2(color, j, v, 0) {
			is_bipartite = false
		}
	}

	fmt.Println("is_bipartite, color =>", is_bipartite, color)

	// トポロジカルソート
	// 与えられた有効グラフに対して、各頂点の向きに沿うように順序付けて並び替えること
	// つまり、タスク間に依存関係がある場合、どんな順番でタスクをこなせば良いかを決めるのに役立つアルゴリズム
	// トポロジカルソートができるグラフは決まりがあって、「ある頂点から出発して、その頂点に戻ってくるような路がない有向グラフ（有向非巡回グラフ（directed acyclic graph, DAG））であること」
	t := [][]int{
		{5},
		{3, 6},
		{5, 7},
		{0, 7},
		{1, 2, 6},
		{},
		{7},
		{0},
	}

	seen := make([]bool, len(t))
	order := make([]int, 0, len(t))

	for v := 0; v < len(t); v++ {
		if seen[v] {
			continue
		}
		rec(t, v, seen)
	}
	fmt.Println("order =>", order)
}

func rec(g [][]int, v int, seen []bool) {
	seen[v] = true
	for _, next_v := range g[v] {
		if seen[next_v] {
			continue
		}
		rec(g, next_v, seen)
	}

	// v-outを記録
	fmt.Println("v-out", v)
}

func dfs2(color []int, g [][]int, v int, cur int) bool {
	color[v] = cur

	for _, next_v := range g[v] {
		fmt.Println("color, cur, v ->", color, cur, v)
		// 隣接頂点がすでに色が確定していた場合
		if color[next_v] != -1 {
			// 同じ色が隣接した場合は2部グラフではない。
			if color[next_v] == cur {
				return false
			}
			// 色が確定した場合には探索しない。
			continue
		}

		// 未確定の場合
		if !dfs2(color, g, next_v, 1-cur) {
			return false
		}
	}
	return true
}

func bfs(g [][]int, s int) {
	n := len(g)
	dist := make([]int, n)
	for i, _ := range dist {
		dist[i] = -1
	}

	q := &Queue{0, 0, make([]int, n)}
	// 開始位置の頂点は処理済とする
	dist[0] = s
	q.push(s)

	for !q.empty() {
		v := q.pop()

		for _, x := range g[v] {
			if dist[x] != -1 {
				continue
			}

			dist[x] = dist[v] + 1
			q.push(x)
		}
		fmt.Println("dist", dist)
	}

	fmt.Println("dist", dist)
	fmt.Println("Queue ->", q.queue)
}

// グラフ探索（幅優先）
func hsearch(g [][]int, s int) {
	bfs(g, s)
}

func dfs(g [][]int, v int, seen []bool) {
	seen[v] = true
	fmt.Println(seen)

	for _, next_v := range g[v] {
		if seen[next_v] {
			continue
		}
		dfs(g, next_v, seen)
	}
}

// グラフ探索（深さ優先）
func dsearch(g [][]int, s int) {
	n := len(g)
	seen := make([]bool, n)

	for v := 0; v < n; v++ {
		if seen[v] {
			continue
		}
		dfs(g, v, seen)
	}
}

// グラフ探索
func search(g [][]int, s int) {
	n := len(g)
	seen := make([]bool, n)

	seen[s] = true
	// queue
	q := &Queue{0, 0, make([]int, n)}
	q.push(s)

	for !q.empty() {
		v := q.pop()

		for _, v := range g[v] {
			if seen[v] {
				continue
			}

			seen[v] = true
			q.push(v)
		}
	}

	fmt.Println(q.queue)
}

type Queue struct {
	tail  int
	head  int
	queue []int
}

func (q *Queue) push(x int) {
	q.queue[q.tail] = x
	q.tail++
}

func (q *Queue) pop() int {
	if q.empty() {
		fmt.Println("error: queue is empty.")
		return -1
	}

	res := q.queue[q.head]
	q.head++
	return res
}

func (q *Queue) empty() bool {
	return q.head == q.tail
}
