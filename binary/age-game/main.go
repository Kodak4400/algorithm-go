package main

import "fmt"

func main() {
	// 年齢当てゲームの実装
	var ans string

	left := 20
	right := 36

	for right-left > 1 {
		mid := left + (right-left)/2

		fmt.Printf("Is the age less than %d (yes / no)\n", mid)

		fmt.Scan(&ans)
		fmt.Println("ans =>", ans)

		if ans == "yes" {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Printf("The age is %d !\n", left)

}
