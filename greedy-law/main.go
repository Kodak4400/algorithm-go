package main

import (
	"fmt"
	"sort"
)

var value = [6]int{500, 100, 50, 10, 5, 1}

func main() {
	// 貪欲法
	// 支払う金額
	x := 1000 // 1 + 2 + 5 + 3 + 2 = 13(500 + 200 + 250 + 30 + 20)
	// 枚数制限
	a := [6]int{
		1, 2, 5, 3, 2, 1,
	}

	result := 0
	for i := 0; i < 6; i++ {
		add := x / value[i]

		if add > a[i] {
			add = a[i]
		}

		x -= value[i] * add
		result += add
	}

	fmt.Println(result)

	// 区間スケジューリング問題
	type pair struct {
		first  int
		second int
	}
	s := []pair{
		{first: 0, second: 2}, //
		{first: 1, second: 3},
		{first: 4, second: 5}, //
		{first: 3, second: 4},
		{first: 6, second: 7}, //
	}

	sort.SliceStable(s, func(i, j int) bool { return s[i].second < s[j].second })

	// 貪欲に選ぶ
	res := 0
	current_end_time := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i].first < current_end_time {
			continue
		}

		res++
		current_end_time = s[i].second
	}
	fmt.Println("res", res)

	// AtCoder Grand Contest 009 A - Multiple Array
	n1 := 5
	a1 := [5]int{0, 1, 3, 5, 4}
	b1 := [5]int{1, 2, 5, 5, 4}

	sum := 0
	for i := n1 - 1; i >= 0; i-- {
		a1[i] += sum // 前回までの操作回数を足す
		amari := a1[i] % b1[i]
		d1 := 0
		if amari != 0 {
			d1 = b1[i] - amari
			sum += d1 // （倍数にしなくてはならないので）余りが操作回数になる
		}
	}
	fmt.Println("sum", sum)
}
