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

	L := make([]int, N)
	for i := 0; i < N; i++ {
		_, l_err := fmt.Scan(&L[i])
		if l_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	// 1 <= i < j < k <= N
	// 三角形の成立条件
	// a + b > c
	// b + c > a
	// c + a > b

	sort.Slice(L, func(i, j int) bool { return L[i] < L[j] })

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				if L[k] != L[j] && L[i] != L[j] && L[k]+L[j] > L[i] {
					fmt.Printf("%d, %d, %d\n", L[i], L[j], L[k])
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
