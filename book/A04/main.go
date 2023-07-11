package main

import (
	"bufio"
	"fmt"
	"math"
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

	ans := make([]int, 10)

	for i := 9; i >= 0; i-- {
		ans[i] = N / int(math.Pow(2, float64(9-i))) % 2
	}

	for _, a := range ans {
		fmt.Printf("%d", a)
	}
}
