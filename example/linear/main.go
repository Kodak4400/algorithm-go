package main

import (
	"fmt"
)

func main() {
	fmt.Println("N個の整数 a0,a1,..., aN-1 と整数vが与えられる。ai = V となるデータが存在するかどうかを判定してください。")

	var N int
	fmt.Print("N:")
	_, n_err := fmt.Scan(&N)
	if n_err != nil {
		fmt.Println("input error.")
		return
	}

	a := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Printf("a[%d]:", i)
		_, err := fmt.Scan(&a[i])
		if err != nil {
			fmt.Println("input error.")
			return
		}
	}

	var V int
	fmt.Print("V:")
	_, v_err := fmt.Scan(&V)
	if v_err != nil {
		fmt.Println("input error.")
		return
	}

	exist := false
	for _, v := range a {
		if v == V {
			exist = true
			break
		}
	}

	fmt.Println(exist)

}
