package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r, g, b int
	_, err := fmt.Scan(&r, &g, &b)
	if err != nil {
		fmt.Println("[Error] input.")
		return
	}

	s := fmt.Sprintf("%d%d%d", r, g, b)
	i, _ := strconv.Atoi(s)
	fmt.Println(i)

	sum := r*100 + g*10 + b
	if sum%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
