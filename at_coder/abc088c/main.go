package main

import (
	"fmt"
)

func main() {
	// 1 0 1
	// 2 1 2
	// 1 0 1

	//   x-0 x x-0
	// 0| 1 0 1
	//  | 2 1 2
	//  | 1 0 1

	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			_, ij_err := fmt.Scan(&c[i][j])
			if ij_err != nil {
				fmt.Println("[Error] input.")
				return
			}
		}
	}

	fmt.Println(c)

	x := make([]int, 3)
	y := make([]int, 3)

	x[0] = 0
	for i := 0; i < 3; i++ {
		y[i] = c[0][i] - x[0]
	}
	for i := 0; i < 3; i++ {
		x[i] = c[i][0] - y[0]
	}

	good := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if x[i]+y[j] != c[i][j] {
				good = false
			}
		}
	}
	if good {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
