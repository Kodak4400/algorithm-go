package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	N := scanInt()

	A := make([]string, N)
	for i, _ := range A {
		A[i] = scanString()
	}

	B := make([]string, N)
	t := ""
	for i, v := range A {
		if i == 0 {
			// 一番上
			t = v[N-1:]
			B[i] = A[i+1][:1] + v[:N-1]
		} else if i < N-1 {
			// 真ん中
			B[i] = A[i+1][:1] + v[1:N-1] + t
			t = A[i][N-1:]
		} else {
			// 一番下
			B[i] = v[1:] + t
		}
	}

	for _, v := range B {
		fmt.Println(v)
	}
}
