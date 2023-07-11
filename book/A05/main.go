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

	ans := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			c := K - (i + j)
			if c <= N && c > 0 {
				ans++
			}
			// for t := 1; t <= N; t++ {
			// 	if i+j+t == K {
			// 		ans++
			// 	}
			// }
		}
	}

	fmt.Println(ans)
}
