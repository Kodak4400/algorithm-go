package main

import (
	"fmt"
)

func ABS(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	// 4
	// x y h
	// 2 3 5
	// 2 1 5
	// 1 2 5
	// 3 2 5

	// マンハッタン距離　・・・ 2つの座標の差の総和を表したもの
	// 高さもあれば、高さも足す。
	// max(H - |X-Cx|-|Y-Cy|, 0)

	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	x := make([]int, N)
	y := make([]int, N)
	h := make([]int, N)

	for i := 0; i < N; i++ {
		_, xyh_err := fmt.Scan(&x[i], &y[i], &h[i])
		if xyh_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	MAX := 100
	// var X = make([][]int, MAX)
	// for i := 0; i < MAX; i++ {
	// 	X[i] = make([]int, MAX)
	// }

	for posY := 0; posY <= MAX; posY++ {
		for posX := 0; posX <= MAX; posX++ {
			// 頂上がどれくらいの高さであってほしいのか。
			// -1はまだよくわからない。0以上は確定しているとき。
			// -2はダメだってわかった時。
			needH := -1

			for i := 0; i < N; i++ {
				if h[i] > 0 {
					// この頂点から見て、頂上がposY, posXの時に,
					// どのくらいの高さがあってほしいか（マンハッタン距離を求める）
					tmp := h[i] + ABS(posY, y[i]) + ABS(posX, x[i])
					if needH == -1 {
						needH = tmp
					} else {
						if needH != tmp {
							needH = -2
							break
						}
					}
				}
			}
			// ダメなら別の場所を探す
			if needH == -2 {
				continue
			}

			for i := 0; i < N; i++ {
				// 高さが0の場合に矛盾が生じないかを確認する。
				// つまり、高さを考えずにマンハッタン距離を求める。
				if h[i] == 0 {
					dist := ABS(posY, y[i]) + ABS(posX, x[i])
					if needH > dist {
						needH = -2
						break
					}
				}
			}
			if needH == -2 {
				continue
			}

			fmt.Println(posX, posY, needH)
		}
	}
}
