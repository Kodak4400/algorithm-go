package main

import (
	"fmt"
)

func main() {
	// 線形探索
	INF := 20000000
	v := 4
	var arr [5]int = [5]int{4, 3, 12, 7, 11}

	exist := false
	found_id := -1
	// [1] 見つけた後は、break
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			exist = true
			found_id = i
			break
		}
	}

	// [2] 最小値を見つける
	min_value := INF
	for i := 0; i < len(arr); i++ {
		if arr[i] < min_value {
			min_value = arr[i]
		}
	}

	// [3] ペア和の最小値を見つける(k以上の範囲)
	k := 10
	pair_min_value := INF
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum <= k {
				continue
			}
			if sum < pair_min_value {
				pair_min_value = sum
			}
		}
	}

	// [4] 部分和問題に対するビットを用いる全探索解法
	// 配列の中から何個かの整数を選んで、総和をwにする。
	// 1 << 5 : 0001 => 00100000(32)
	// 1 & 1 << 5 :  00000001 & 00100000 = 00000000(false)
	// 32 & 1 << 5 :  00100000 & 00100000 = 00100000(true)
	w := 26
	sum_exist := false
	var sum_slice []int
	for bit := 0; bit < (1 << len(arr)); bit++ {
		sum := 0
		sum_slice = nil
		for i := 0; i < len(arr); i++ {
			if bit&(1<<i) != 0 {
				sum += arr[i]
				sum_slice = append(sum_slice, arr[i])
			}
		}
		if sum == w {
			sum_exist = true
			break
		}
	}

	// [5] 部分和を再帰関数を用いる全探索
	// N = 4
	// a=(3, 2, 6, 5)
	// W = 14
	a := [4]int{3, 2, 6, 5}
	r := p(4, 14, a)
	fmt.Println("14になった？ =>", r)

	fmt.Println("線形探索(存在有無) =>", exist)
	fmt.Println("線形探索(添字取得) =>", found_id)
	fmt.Println("線形探索(最小値取得) =>", min_value)
	fmt.Println("線形探索(ペア和の最小値取得) =>", pair_min_value)
	fmt.Printf("線形探索(部分和問題に対するビットを用いる全探索解法) => sum_exist: %t, w: %d sum_slice: %d \n", sum_exist, w, sum_slice)
}

func p(i int, w int, a [4]int) bool {
	if i == 0 {
		if w == 0 {
			fmt.Printf("i=>%d, w=>%d, a=>%v\n", i, w, a)
			return true
		} else {
			return false
		}
	}

	if p(i-1, w, a) {
		fmt.Printf("i=>%d, w=>%d, a=>%v\n", i, w, a)
		return true
	}

	if p(i-1, w-a[i-1], a) {
		fmt.Printf("i=>%d, w=>%d, a=>%v\n", i, w, a)
		return true
	}

	return false
}
