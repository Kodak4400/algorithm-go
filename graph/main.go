package main

import "fmt"

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
