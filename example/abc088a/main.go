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

	var A int
	_, a_err := fmt.Scan(&A)
	if a_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	if (N % 500) <= A {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
