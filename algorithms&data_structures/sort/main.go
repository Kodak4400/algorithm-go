package main

import "fmt"

func main() {
	// 挿入ソート
	//  - in-placeなソート
	//  - 安定ソート
	a := []int{4, 1, 3, 5, 2}
	fmt.Println("insertion sort before ->", a)
	insertionSort(a)
	fmt.Println("insertion sort after  ->", a)

	// マージソート
	//  - in-placeなソートではない
	//  - 安定ソート
	b := []int{12, 9, 15, 3, 8, 17, 6, 1}
	fmt.Println("merge sort before ->", b)
	mergeSort(b, 0, 8)
	fmt.Println("merge sort after  ->", b)

	// クイックソート
	//  - in-placeなソートではない
	//  - 安定ソートではない
	c := []int{10, 12, 15, 3, 8, 17, 4, 1}
	fmt.Println("quick sort before ->", c)
	quickSort(c, 0, 8)
	fmt.Println("quick sort after  ->", c)

	// ヒープソート
	//  - in-placeなソート
	//  - 安定ソートではない
	d := []int{10, 12, 15, 3, 8, 17, 4, 1}
	fmt.Println("heap sort before ->", d)
	heapSort(d)
	fmt.Println("heap sort after  ->", d)

	// バケットソート
	//  - in-placeなソートではない
	//  - 安定ソート
	e := []int{1, 3, 2, 1, 2, 1}
	fmt.Println("bucket sort before ->", e)
	bucketSort(e)
	fmt.Println("bucket sort after  ->", e)

}

func bucketSort(e []int) {
	n := len(e)

	// 各要素の個数をカウントする
	num := make([]int, n)
	for i := 0; i < n; i++ {
		num[e[i]]++
	}
	fmt.Println("バケットのカウント数", num) // [0 3 2 1 0 0]

	// numの累積和をとる
	sum := make([]int, n)
	for v := 1; v < n; v++ {
		sum[v] = sum[v-1] + num[v]
	}
	fmt.Println("バケットのカウントの累積和", sum) //  [0 3(1,1,1) 5(2,2) 6(3) 6 6]

	// sumをもとにソート
	e2 := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		e2[sum[e[i]]-1] = e[i] // e2[2]=1, e2[4]=2, e2[1]=1, e2[3]=2, e2[5]=3, e2[0]=1
		sum[e[i]]--
	}
	fmt.Println("e2", e2)
	// e = e2
}

func heapify(d []int, i int, n int) {
	child1 := i*2 + 1
	if child1 >= n {
		return
	}

	if child1+1 < n && d[child1+1] > d[child1] {
		child1++
	}

	if d[child1] <= d[i] {
		return
	}

	swap(d, i, child1)
	heapify(d, child1, n)
}

func heapSort(d []int) {
	n := len(d)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(d, i, n)
	}
	for i := n - 1; i > 0; i-- {
		swap(d, 0, i)
		heapify(d, 0, i)
	}
}

func quickSort(c []int, left int, right int) {
	if right-left <= 1 {
		return
	}

	pivotIndex := (left + right) / 2 // 適当に中心点を決める
	pivot := c[pivotIndex]
	swap(c, pivotIndex, right-1) // pivotと右端をswap

	i := left // iは左詰めされたpivot未満要素の右端を表す
	for j := left; j < right-1; j++ {
		if c[j] < pivot {
			swap(c, i, j)
			i++
		}
	}
	swap(c, i, right-1) // pivotを適切な場所へ挿入

	quickSort(c, left, i)
	quickSort(c, i+1, right)
}

func swap(c []int, i int, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

func mergeSort(b []int, left int, right int) {
	if right-left == 1 {
		return
	}
	mid := left + (right-left)/2

	mergeSort(b, left, mid)
	mergeSort(b, mid, right)

	buf := make([]int, 0, len(b))
	for i := left; i < mid; i++ {
		buf = append(buf, b[i])
	}
	for i := right - 1; i >= mid; i-- {
		buf = append(buf, b[i])
	}

	// 併合する
	indexLeft := 0
	indexRight := len(buf) - 1
	for i := left; i < right; i++ {
		if buf[indexLeft] <= buf[indexRight] {
			b[i] = buf[indexLeft]
			indexLeft++
		} else {
			b[i] = buf[indexRight]
			indexRight--
		}
	}
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		v := a[i] // 挿入したい値

		// vを挿入する適切な場所jを探す
		j := i
		for j > 0 {
			if a[j-1] > v {
				a[j] = a[j-1]
			} else {
				break
			}
			a[j-1] = v
			j--
		}
	}
}
