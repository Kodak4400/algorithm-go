package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

type person struct {
	id, a, b int
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	N := scanInt()
	p := make([]person, N)
	for i, _ := range p {
		p[i].id = i + 1
		p[i].a, p[i].b = scanInt(), scanInt()
	}

	sort.Slice(p, func(i, j int) bool {
		n := p[i].a*(p[j].a+p[j].b) - p[j].a*(p[i].a+p[i].b)
		if n == 0 {
			return p[i].id < p[j].id
		}
		return n > 0
	})

	ans := make([]string, 0, N)
	for _, v := range p {
		ans = append(ans, strconv.Itoa(v.id))
	}
	fmt.Println(strings.Join(ans, " "))
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
