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
	search(g, 1)
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
		v := q.head

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

func (q *Queue) empty() bool {
	return q.head == q.tail
}
