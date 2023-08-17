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

	res := false
	// if N/11 != 0 {
	// 	res = true
	// } else if N/7 != 0 {
	// 	res = true
	// } else if N/4 != 0 {
	// 	res = true
	// }

	// if res {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }

	CMAX := 100 / 4
	DMAX := 100 / 7

	for i := 0; i <= CMAX; i++ {
		for j := 0; j <= DMAX; j++ {
			if i*4+j*7 == N {
				res = true
			}
		}
	}

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
