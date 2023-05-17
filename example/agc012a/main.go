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

	t := 3 * N

	a := make([]int, t)
	for i := 0; i < t; i++ {
		_, a_err := fmt.Scan(&a[i])
		if a_err != nil {
			fmt.Println("[Error] input.")
			return
		}
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	// 6 5 4 3 2 1
	// 5 4 3 2 1 0
	sum := 0
	j := 0
	for i := t - 2; i > 0; i = i - 2 {
		if j == N {
			break
		}
		sum += a[i]
		j++
		fmt.Println(i)
	}

	fmt.Println(sum)
}
