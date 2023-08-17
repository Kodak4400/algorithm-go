package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	_, n_err := fmt.Scan(&S)
	if n_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	fmt.Println(700 + 100*strings.Count(S, "o"))
}
