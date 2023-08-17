package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, x int
	_, n_err := fmt.Scan(&N, &x)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	a := make([]int, N)
	for i := 0; i < N; i++ {
		_, a_err := fmt.Scan(&a[i])
		if a_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	b := make([]int, N, N)

	ans := 0
	if x < a[0] {
		fmt.Println(ans)
		return
	}

	for true {
		if x == 0 {
			break
		}
		for i := 0; i < N; i++ {
			if x-a[i] >= 0 {
				x -= a[i]
				b[i] += a[i]
				ans++
				if x == 0 {
					break
				}
			}
		}
	}
	fmt.Println(ans)
	fmt.Println(b)
}
