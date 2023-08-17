package main

import (
	"fmt"
)

func main() {
	// 10000 <= A <= B <= 99999
	var A, B int
	_, n_err := fmt.Scan(&A, &B)
	if n_err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	ans := 0
	for i := A; i < B; i++ {
		s := i % 10 // 1
		t := i / 10000 % 10
		u := i / 10 % 10
		v := i / 1000 % 10
		if s == t && u == v {
			ans++
		}
	}
	fmt.Println(ans)
}
