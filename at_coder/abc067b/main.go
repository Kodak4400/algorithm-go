package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	_, n_err := fmt.Scan(&N, &K)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	l := make([]int, N)
	for i := 0; i < N; i++ {
		_, l_err := fmt.Scan(&l[i])
		if l_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })

	sum := 0
	for i := N - K; i < N; i++ {
		sum += l[i]
	}

	fmt.Println(sum)
}
