package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	A := make([]int, N, N)
	for i := 0; i < N; i++ {
		_, n_err := fmt.Scan(&A[i])
		if n_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	fmt.Println(A[len(A)-1] - A[0])
}
