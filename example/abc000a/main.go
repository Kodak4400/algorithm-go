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
	for i := 0; i < N; i++ {
		A[i] = scanString()
	}

	var r string
	for i, a := range A {
		if i == 0 {
			r = a[N-1:]
			A[i] = A[i+1][:1] + a[:N-1]
		} else if i == N-1 {
			A[i] = a[1:] + r
		} else {
			t := a[N-1:]
			A[i] = A[i+1][:1] + a[1:N-1] + r
			r = t
		}
	}

	for _, a := range A {
		fmt.Println(a)
	}
}
