package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N, M := scanInt(), scanInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	B := make([]int, M)
	for i := 0; i < M; i++ {
		B[i] = scanInt()
	}

	sort.Ints(A)
	sort.Ints(B)

	start := 0
	end := int(1e9) + 1
	for {
		// fmt.Println(start, end)

		v := (start + end) / 2
		sell := sort.Search(N, func(i int) bool { return A[i] > v })
		buy := M - sort.Search(M, func(i int) bool { return B[i] >= v })

		// fmt.Println("v, sell, buy ->", v, sell, buy)

		if start == end {
			// fmt.Println(sell, buy)
			break
		}

		if sell < buy {
			start = v + 1
		} else if buy <= sell {
			end = v
		}
	}
	fmt.Println(start)
}
