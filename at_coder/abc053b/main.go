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
	cnt := 0
	res := 0
	for i, _ := range ss {
		if string(ss[i]) == "A" {
			cnt++
			continue
		}
		if cnt > 0 {
			cnt++
		}
		if string(ss[i]) == "Z" {
			if cnt > res {
				res = cnt
			}
		}
	}
	fmt.Println(res)
}
