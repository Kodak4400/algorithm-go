package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("0と1のみから成る3桁の番号sが与えられます。1が何個含まれるかを求めてください。")
	fmt.Print("INPUT:")

	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	fmt.Println(strings.Count(s, "1"))
}
