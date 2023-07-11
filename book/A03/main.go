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

	N, X := scanInt(), scanInt()

	A := make([]int, N)
	for i, _ := range A {
		A[i] = scanInt()
	}

	ans := false
	for _, a := range A {
		if a == X {
			ans = true
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
