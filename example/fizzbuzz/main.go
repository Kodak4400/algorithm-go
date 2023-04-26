package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(i)
	}
}

func main() {
	fmt.Println("1から標準入力から入れる値(INPUT値)までの数を順に出力するプログラム。\nただし、3の倍数のときは数の代わりに「Fizz」5の倍数のときは「Buzz」3と5の両方の倍数の場合には「FizzBuzz」を出力する。")
	fmt.Print("INPUT:")

	var n int
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("err: INPUT.")
		return
	}

	for i := 1; i <= n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
