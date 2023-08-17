package main

import (
	"fmt"
)

// N := 8
var a = [8]int{3, 5, 8, 10, 14, 17, 21, 39}

func main() {
	// 配列から目的の値を探索する二分探索法
	fmt.Println("binary_search1(10) ->", binary_search1(10))
	fmt.Println("binary_search1(3) ->", binary_search1(3))
	fmt.Println("binary_search1(39) ->", binary_search1(39))
	fmt.Println("binary_search1(-100) ->", binary_search1(-100))
	fmt.Println("binary_search1(9) ->", binary_search1(9))
	fmt.Println("binary_search1(100) ->", binary_search1(100))

	// C++ std::lower_bound()
	// ソート済み配列aにおいて、a[i] >= keyという条件を満たす最小の添字iを返す(正確にはiteratorを返す)
	// 処理に要する計算量は、配列サイズをNとしてO(log N)
	// - 配列aの中に値keyがなくても、key以上の値の範囲で最小値がわかる
	// - 配列aの中に値keyが複数あったとき、そのうちの最小の添字がわかる

	fmt.Println("binary_search2() ->", binary_search2())

	// 二分探索法を用いて「ペア和を最適化する問題」に対する全探索解放を高速化する
}

func binary_search1(key int) int {
	result := -1
	left := 0
	right := len(a) - 1

	for right >= left {
		mid := left + (right-left)/2
		if a[mid] == key {
			result = mid
			break
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return result
}

// 一般化した二分探索法の基本形
func p(x int) bool {
	var r = false
	if a[x] == 21 {
		r = true
	}
	return r
}

func binary_search2() int {
	var mid = 0
	left := 0
	right := len(a) - 1

	for right-left > 1 {
		// 0 + (7-0)/2 = 3
		// 3 + (7-3)/2 = 5
		// 5 + (7-5)/2 = 6
		mid = left + (right-left)/2
		if p(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
