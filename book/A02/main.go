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

	N, K := scanInt(), scanInt()

	P := make([]int, N)
	Q := make([]int, N)
	for i, _ := range P {
		P[i], Q[i] = scanInt(), scanInt()
	}

	ans := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if P[i]+Q[j] == K {
				ans = true
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
