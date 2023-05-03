package main

import (
	"fmt"
)

func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println("1以上N以下の整数のうち、10進法で各桁の和がA以上B以下であるものについて、総和を求めてください")
	fmt.Print("INPUT:")

	var N int
	var A int
	var B int
	_, err := fmt.Scan(&N, &A, &B)
	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	total := 0
	for i := 0; i <= N; i++ {
		var sum = findSumOfDigits(i)
		if sum >= A && sum <= B {
			total += i
		}
	}
	fmt.Println(total)
}
