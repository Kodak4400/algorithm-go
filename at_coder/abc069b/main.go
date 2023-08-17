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

	s := strings.Split(S, "")
	fmt.Printf("%s%d%s\n", s[0], len(s)-2, s[len(s)-1])
}
