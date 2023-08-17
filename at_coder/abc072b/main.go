package main

import (
	"fmt"
)

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("[error] input.")
		return
	}

	ss := []rune(s)
	res := ""
	for i := 0; i < len(s); i++ {
		if (i+1)%2 != 0 {
			// fmt.Println(string(ss[i]))
			res += string(ss[i])
		}
	}
	fmt.Println(res)
}
