package main

import (
	"fmt"
)

func main() {
	fmt.Println("10000 円札と、5000 円札と、1000 円札が合計で N枚あって、合計金額が Y円であったという。\nこのような条件を満たす各金額の札の枚数の組を1つ求めなさい。そのような状況が存在し得ない場合には -1 -1 -1 と出力しなさい。")

	var N int
	fmt.Print("N:")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	var Y int
	fmt.Print("Y:")
	_, y_err := fmt.Scan(&Y)
	if y_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	res10000, res5000, res1000 := -1, -1, -1

	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {
			c := N - a - b
			total := 10000*a + 5000*b + 1000*c
			if total == Y {
				res10000 = a
				res5000 = b
				res1000 = c
			}
		}
	}
	fmt.Println(res10000, res5000, res1000)

}
