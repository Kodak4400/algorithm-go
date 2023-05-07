package main

import (
	"fmt"
)

func main() {
	fmt.Println("N個の整数d[0],d[1]...d[N-1]が与えられます\nこの中に何種類異なる値があるでしょうか？")

	var N int
	fmt.Print("N:")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	d := make([]int, N, N)
	fmt.Print("d:")
	for i := 0; i < N; i++ {
		_, a_err := fmt.Scan(&d[i])
		if a_err != nil {
			fmt.Println("err: INPUT.")
			return
		}
	}

	// バケット法
	num := [110]int{0} // 1 <= d[i] <= 100までという前提あり
	fmt.Println("num =>", num)
	for i := 0; i < N; i++ {
		num[d[i]]++
	}

	res := 0
	for i := 0; i < 110; i++ {
		if num[i] > 0 {
			res++
		}
	}

	fmt.Println("num =>", num)
	fmt.Println("res =>", res)

	// ハッシュを使う
	obj := make(map[int]int)
	for i := 0; i < N; i++ {
		obj[d[i]] = d[i]
	}

	fmt.Println("obj =>", len(obj))
}
