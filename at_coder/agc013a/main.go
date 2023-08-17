package main

import (
	"fmt"
)

func main() {
	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	A := make([]int, N)
	for i := 0; i < N; i++ {
		_, a_err := fmt.Scan(&A[i])
		if a_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	// 単調減少 => xが増えればyが減るようなこと
	// 単調非減少 => xが増えてもyが減らない。

	// flgが-1は、よくわからない。
	// 0なら単調非増加。1なら単調非減少。
	flg := -1

	i := 0
	if A[i] < A[i+1] {
		flg = 0
	} else if A[i] > A[i+1] {
		flg = 1
	}

	// for i := 0; i < N; i++ {
	// 	if A[i] < A[i+1] {
	// 		flg = 0
	// 	}
	// 	if A[i] > A[i+1] {
	// 		flg = 1
	// 	}

	// }
}
