package main

import (
	"fmt"
)

func main() {
	fmt.Println("2つの整数a,bが与えられます。aとbの積が偶数か奇数かを判定してください")
	fmt.Print("INPUT:")

	var a int
	var b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	sum := a * b
	if sum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
