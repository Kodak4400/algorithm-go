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

/*
15 3
62 65 41 13 20 11 18 44 53 12 18 17 14 10 39
4 13
3 10
2 15
*/
func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	N, Q := scanInt(), scanInt()

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = scanInt()
	}

	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		L[i], R[i] = scanInt(), scanInt()
	}

	fmt.Println(A)

	ruiseki := make([]int, N+1)
	ruiseki[0] = 0
	for i := 1; i <= N; i++ {
		ruiseki[i] = ruiseki[i-1] + A[i]
	}

	fmt.Println(ruiseki)
	for i := 0; i < Q; i++ {
		fmt.Println(ruiseki[R[i]], ruiseki[L[i]-1])
		fmt.Println(ruiseki[R[i]] - ruiseki[L[i]-1])
	}

}
