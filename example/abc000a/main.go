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

// type persion struct {
// 	id, a, b int
// }

// func main() {
// 	rd := bufio.NewReader(os.Stdin)
// 	wr := bufio.NewWriter(os.Stdout)
// 	defer wr.Flush()

// 	var n int
// 	fmt.Fscan(rd, &n)

// 	p := make([]persion, n)
// 	for i := range p {
// 		p[i].id = i + 1
// 		fmt.Fscan(rd, &p[i].a, &p[i].b)
// 	}

// 	fmt.Println(p)
// 	sort.Slice(p, func(i, j int) bool {
// 		d := p[i].a*(p[j].a+p[j].b) - p[j].a*(p[i].a+p[i].b)
// 		if d == 0 {
// 			return p[i].id < p[j].id
// 		}
// 		return d > 0
// 	})

// 	ans := make([]string, 0, len(p))
// 	fmt.Println(ans)
// 	for i := range p {
// 		ans = append(ans, strconv.Itoa(p[i].id))
// 	}

// 	fmt.Fprintln(wr, strings.Join(ans, " "))
// }

// func main() {
// 	var N int
// 	fmt.Scanf("%d", &N)

// 	A := make([]int, N)
// 	B := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		fmt.Scanf("%d %d", &A[i], &B[i])
// 	}

// 	t := make([]int64, N)
// 	s := make([]int64, N)
// 	for i := 0; i < N; i++ {
// 		h := seikouritu(A[i], B[i])
// 		t[i] = h
// 		s[i] = h
// 	}

// 	sort.SliceStable(t, func(i, j int) bool {
// 		return t[i] > t[j]
// 	})

// 	ans := make([]int, N)
// 	for i, v := range t {
// 		for j := 0; j < N; j++ {
// 			if s[j] == v {
// 				ans[i] = j + 1
// 				s[j] = -1
// 			}
// 		}
// 	}

// 	for _, v := range ans {
// 		fmt.Printf("%d ", v)
// 	}
// }

func seikouritu(a int, b int) int64 {
	return int64(a * (a + b))
}

func getValues(m map[int]float64) []float64 {
	values := []float64{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func getKeys(m map[int]float64) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func insertionSort(a []uint64) {
	for i := 1; i < len(a); i++ {
		v := a[i] // 挿入したい値

		// vを挿入する適切な場所jを探す
		j := i
		for j > 0 {
			if a[j-1] < v {
				a[j] = a[j-1]
			} else {
				break
			}
			a[j-1] = v
			j--
		}
	}
}
