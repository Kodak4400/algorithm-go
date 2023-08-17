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

	fmt.Printf("%d\n", (N+1)/2)
}
