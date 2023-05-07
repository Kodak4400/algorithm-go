package main

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s) // rune型はcode pointを単位として扱う型
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("英小文字からなる文字列Sが与えられます。Tが空文字列である状態から始めて、以下の操作を好きな回数繰り返すことで S=Tとすることができるか判定してください。\n - Tの末尾に \"dream\", \"dreamer\", \"erase\", \"eraser\"のいずれかを追加する。")

	var S string
	fmt.Print("S:")
	_, n_err := fmt.Scan(&S)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	// Greedy アルゴリズム

	divide := []string{"dream", "dreamer", "erase", "eraser"}

	// 後ろから解くために反転
	revS := reverse(S)

	divide2 := make([]string, 4, 4)
	for i := 0; i < 4; i++ {
		divide2[i] = reverse(divide[i])
	}

	can := true
	i := 0
	fmt.Println("revS", revS)
	for i < len(S) {
		can2 := false // 4個の文字のどれかでdivideできるか
		for j := 0; j < 4; j++ {
			var d = divide2[j]
			fmt.Println(i, d)
			if i < len(S) && revS[i:i+len(d)] == d {
				can2 = true
				i += len(d)
			}
		}
		if !can2 {
			can = false
			break
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
