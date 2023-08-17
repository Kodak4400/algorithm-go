package main

import "fmt"

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

func main() {
	// ヒープ
	heap := make([]int, 12, 15)
	for i, v := range []int{19, 12, 15, 10, 7, 6, 1, 3, 7, 5, 3, 2} {
		heap[i] = v
	}

	fmt.Printf("push前のheap =>%v\n", heap)
	push(17, heap)
	fmt.Printf("push後のheap =>%v\n", heap)
	pop(heap)
	fmt.Printf("pop後のheap =>%v\n", heap)

}
