package main

import (
	"fmt"
)

func main() {
	// W H N
	// 5 4 2
	// 2 1 1
	// 3 3 4

	// a = 1 xiより小さい 0~2を引く
	// a = 2 xiより大きい 2~5を引く
	// a = 3 yiより小さい 0~3を引く
	// a = 4 yiより大きい 3~4を引く

	// (5, 4) => 正方形面積 => 5 * 4
	// (5-2, 4-(4-3)) => (3, 3)

	// 5 4 3 (5, 4)
	// 2 1 1 (3, 4)
	// 3 3 4 (3, 3)
	// 1 4 2 (0, 3)

	// 1 3
	// 3

	var W, H, N int
	_, err := fmt.Scan(&W, &H, &N)
	if err != nil {
		fmt.Println("[error] input.")
		return
	}

	x := make([]int, N)
	y := make([]int, N)
	a := make([]int, N)

	for i := 0; i < N; i++ {
		_, xya_err := fmt.Scan(&x[i], &y[i], &a[i])
		if xya_err != nil {
			fmt.Println("[error] input.")
			return
		}
	}

	// A[x][y] => [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
	var A = make([][]int, W)
	for x := 0; x < W; x++ {
		A[x] = make([]int, H)
	}

	for i := 0; i < N; i++ {
		switch a[i] {
		case 1:
			for k := 0; k < x[i]; k++ {
				for j := 0; j < H; j++ {
					A[k][j]++
				}
			}
		case 2:
			for k := x[i]; k < W; k++ {
				for j := 0; j < H; j++ {
					A[k][j]++
				}
			}
		case 3:
			for k := 0; k < W; k++ {
				for j := 0; j < y[i]; j++ {
					A[k][j]++
				}
			}
		case 4:
			for k := 0; k < W; k++ {
				for j := y[i]; j < H; j++ {
					A[k][j]++
				}
			}
		}
	}

	fmt.Println(A)

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if A[i][j] == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)

	// w := W
	// h := H

	// for i := 0; i < N; i++ {
	// 	switch a[i] {
	// 	case 1:
	// 		W -= x[i]
	// 	case 2:
	// 		W -= w - x[i]
	// 	case 3:
	// 		H -= y[i]
	// 	case 4:
	// 		H -= h - y[i]
	// 	}
	// }

	// fmt.Println(W, H)

	// if W < 0 {
	// 	W = 0
	// }
	// if H < 0 {
	// 	H = 0
	// }
	// fmt.Println(W * H)
}
