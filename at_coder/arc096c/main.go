package main

import (
	"fmt"
)

func cmin(ans *int, min int) {
	if *ans > min {
		*ans = min
	}
}

func main() {
	var A, B, C, X, Y int
	_, abcxy_err := fmt.Scan(&A, &B, &C, &X, &Y)
	if abcxy_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	// ABピザを買う全探索
	// AB = 1買った時...
	// AB = 2買った時...

	ans := 5000 * 1000000
	for i := 0; i < 201010; i++ {
		sm := C * i

		x := X - i/2
		y := Y - i/2

		if 0 < x {
			sm += x * A
		}
		if 0 < y {
			sm += y * B
		}

		cmin(&ans, sm)
	}

	fmt.Println(ans)
}
