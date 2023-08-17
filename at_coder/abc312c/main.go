package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n := nextInt()
	m := nextInt()
	a := nextInts(n)
	b := nextInts(m)

	sort.Ints(a)
	sort.Ints(b)

	fmt.Println(a, b)

	start := 0
	end := int(1e9) + 1 /* 1000000000 */
	for {
		fmt.Println("start, end ->", start, end)

		v := (start + end) / 2
		sell := sort.Search(n, func(i int) bool { return a[i] > v })
		buy := m - sort.Search(m, func(i int) bool { return b[i] >= v })

		fmt.Println("v, sell, buy ->", v, sell, buy)

		if start == end {
			// fmt.Println(sell, buy)
			break
		}

		if sell < buy { // price up
			start = v + 1
		} else if buy <= sell {
			end = v
		}
	}
	fmt.Println(start)
}

func init() {
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		if len(os.Args) == 2 {
			fmt.Fprintf(os.Stderr, "filename is required")
			os.Exit(1)
		}
		debug(os.Args[2])
	}

	sc.Split(bufio.ScanWords)

	buf := make([]byte, 10*1024)
	sc.Buffer(buf, math.MaxInt32)
}

func debug(filename string) {
	testFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: no such file", filename)
		os.Exit(1)
	}
	sc = bufio.NewScanner(testFile)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	n, err := strconv.Atoi(next())
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	return n
}

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}
