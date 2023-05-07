package main

import (
	"fmt"
)

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func main() {
	fmt.Println("シカの AtCoDeer くんは二次元平面上で旅行をしようとしています。\nAtCoDeer くんの旅行プランでは、時刻0時に点(0,0)を出発し、1以上N以下の各iに対し、時刻tiに点(xi,yi)を訪れる予定です。\nAtCodeerくんが時刻tに点(x,y)にいる時、時刻t+1には点(x+1,y),(x-1,y),(x,y+1),(x,y-1)のうちいずれかに存在することができます。その場にとどまることはできないことに注意してください。\nAtCodeerくんの旅行プランが実行可能か判定してください。")

	var N int
	fmt.Print("N=")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	t := make([]int, N+1, N+1)
	x := make([]int, N+1, N+1)
	y := make([]int, N+1, N+1)
	t[0], x[0], y[0] = 0, 0, 0
	for i := 0; i < N; i++ {
		fmt.Print("t,x,y=")
		_, txy_err := fmt.Scan(&t[i+1], &x[i+1], &y[i+1])
		if txy_err != nil {
			fmt.Println("err: INPUT.")
			return
		}
	}

	fmt.Println("t,x,y =", t, x, y)

	can := true
	for i := 0; i < N; i++ {
		dt := t[i+1] - t[i]
		dist := abs(x[i+1]-x[i]) + abs(y[i+1]-y[i])
		fmt.Println("dt =", dt)
		fmt.Println("dist =", dist)
		if dt < dist {
			can = false
		}
		if dist%2 != dt%2 {
			can = false
		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
