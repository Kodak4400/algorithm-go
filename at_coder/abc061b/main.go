package main

import (
	"fmt"
)

func main() {
	// 4 3
	// 1, 2
	// 2, 3
	// 1, 4

	// a = {1, 2, 1}
	// b = {2, 3, 4}

	// 1 - 2 - 3
	//   - 4

	// 2 2 1 1
	var N, M int
	_, err := fmt.Scan(&N, &M)
	if err != nil {
		fmt.Println("[error] input.")
		return
	}

	a := make([]int, M)
	b := make([]int, M)

	for i := 0; i < M; i++ {
		_, ab_err := fmt.Scan(&a[i], &b[i])
		if ab_err != nil {
			fmt.Println("[error] input.")
			return
		}
	}

	michi := make([]int, N, N)

	for i := 0; i < M; i++ {
		michi[a[i]-1]++
		michi[b[i]-1]++
	}

	fmt.Println(michi)
}
