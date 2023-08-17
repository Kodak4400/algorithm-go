package main

import (
	"fmt"
)

func main() {
	fmt.Println("黒板にN個の正の整数A1,A2....ANが書かれている\nすぬけ君は黒板に書かれている整数すべて偶数であるとき、次の操作が行える。")
	fmt.Println("黒板に書かれている整数すべてを、2で割ったものに置き換える。")
	fmt.Println("すぬけ君は最大で何回操作を行うことができるかを求めてください。")
	fmt.Print("INPUT:")

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := 0
	var exist_odd bool
	for true {
		exist_odd = false
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				exist_odd = true
			}
		}

		if exist_odd {
			break
		}

		for i := 0; i < n; i++ {
			a[i] /= 2
		}
		res++
	}

	fmt.Println(res)
}
