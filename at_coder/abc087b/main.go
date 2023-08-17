package main

import (
	"fmt"
)

func main() {
	fmt.Println("500円玉をA枚、100円玉をB枚、50円玉をC枚持っています。\nこれらの硬化の中から何枚かを選び合計金額をちょうどX円にする方法は何通りあるか")
	fmt.Print("INPUT:")

	var A int
	var B int
	var C int
	var X int
	_, err := fmt.Scan(&A, &B, &C, &X)
	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	res := 0
	var total int
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				total = 500*a + 100*b + 50*c
				if total == X {
					res++
				}
			}
		}
	}

	fmt.Println(res)
}
