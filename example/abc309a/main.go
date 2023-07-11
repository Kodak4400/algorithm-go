package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	A, B := scanInt(), scanInt()

	f := false

	C := B - A
	if C == 1 && A%3 != 0 {
		f = true
	}

	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
