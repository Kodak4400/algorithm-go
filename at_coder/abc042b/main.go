package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, L int
	_, n_err := fmt.Scan(&N, &L)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	S := make([]string, N)
	for i := 0; i < N; i++ {
		_, l_err := fmt.Scan(&S[i])
		if l_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })

	sum := ""
	for i := 0; i < N; i++ {
		sum += S[i]
	}

	fmt.Println(sum)
}
