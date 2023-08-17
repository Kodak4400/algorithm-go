package main

import (
	"fmt"
	"strconv"
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

	ss := ""
	for i, _ := range s {
		if i == 3 {
			ss += "8"
			continue
		}
		ss += s[i]
	}
	fmt.Println(ss)

	split := strings.Split(S, "/")
	year, _ := strconv.Atoi(split[0])

	fmt.Printf("%d/%s/%s\n", year+1, split[1], split[2])
}
