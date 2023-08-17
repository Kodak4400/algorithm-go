package main

import (
	"fmt"
	"math"
)

func main() {
	var N, X int
	_, err := fmt.Scan(&N, &X)
	if err != nil {
		fmt.Println("[error] input.")
		return
	}

	m := make([]int, N)
	for i := 0; i < N; i++ {
		_, m_err := fmt.Scan(&m[i])
		if m_err != nil {
			fmt.Println("[error] input.")
			return
		}
	}

	res := 0
	min := math.MaxInt
	sum := 0
	for i, _ := range m {
		if min > m[i] {
			min = m[i]
		}
		sum += m[i]
		res++
	}

	res += (X - sum) / min

	fmt.Println(res)
}
