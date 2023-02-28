package main

import (
	"fmt"
)

var memo = [7]int{-1, -1, -1, -1, -1, -1, -1}

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

	fmt.Println("線形探索(存在有無) =>", exist)
	fmt.Println("線形探索(添字取得) =>", found_id)
	fmt.Println("線形探索(最小値取得) =>", min_value)
	fmt.Println("線形探索(ペア和の最小値取得) =>", pair_min_value)
	fmt.Printf("線形探索(部分和問題に対するビットを用いる全探索解法) => sum_exist: %t, w: %d sum_slice: %d \n", sum_exist, w, sum_slice)

	// 再帰と分割統治法
	p(5)

	// ユークリッドの互除法で最大公約数を求める
	fmt.Println("gcd(51, 15) =>", gcd(51, 15))
	fmt.Println("gcd(15, 51) =>", gcd(15, 51))

	// フィボナッチ数列を求める
	fmt.Println("fibo(6) =>", fibo(6))

	// フィボナッチ数列を求める(for版)
	fibofor := [7]int{0, 1}
	for n := 2; n < 7; n++ {
		fibofor[n] = fibofor[n-1] + fibofor[n-2]
		fmt.Printf("%d 項目 = %d\n", n, fibofor[n])
	}

	// フィボナッチ数列を求める(メモ化)
	fmt.Println("mfibo(6) =>", mfibo(6))
}

// ------------------------

func p(n int) int {
	fmt.Printf("func p(%d) を呼びました\n", n)

	if n == 0 {
		return 0
	}

	result := n + p(n-1)
	fmt.Printf("%d までの和 = %d\n", n, result)
	return result
}

func gcd(m int, n int) int {
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}

func fibo(n int) int {
	fmt.Printf("func fibo(%d) を呼びました\n", n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := fibo(n-1) + fibo(n-2)
	fmt.Printf("%d 項目 = %d\n", n, result)

	return result
}

func mfibo(n int) int {
	fmt.Printf("func mfibo(%d) を呼びました\n", n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = mfibo(n-1) + mfibo(n-2)

	return memo[n]
}
