package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	var S string
	fmt.Fscan(in, &S)
	ss := strings.Split(S, "")

	mp := map[int][]int{}
	for i := 0; i < N; i++ {
		var c int
		fmt.Fscan(in, &c)
		mp[c] = append(mp[c], i)
	}

	for i := 1; i <= M; i++ {
		if len(mp[i]) <= 1 {
			continue
		}
		last := ss[mp[i][len(mp[i])-1]]
		for j := len(mp[i]) - 1; j > 0; j-- {
			ss[mp[i][j]] = ss[mp[i][j-1]]
		}
		ss[mp[i][0]] = last
	}
	fmt.Fprintln(out, strings.Join(ss, ""))
}

// func main() {
// 	sc = bufio.NewScanner(os.Stdin)
// 	sc.Split(bufio.ScanWords)

// 	// 200000
// 	N, M := scanInt(), scanInt()

// 	S := scanString()
// 	ss := strings.Split(S, "")

// 	C := make([]int, N)

// 	t := make([][]string, M)
// 	for j := 0; j < M; j++ {
// 		t[j] = make([]string, 0, N)
// 	}

// 	for i := 0; i < N; i++ {
// 		C[i] = scanInt()

// 		for j := 0; j < M; j++ {
// 			if C[i] == j+1 {
// 				t[j] = append(t[j], ss[i])
// 			}
// 		}
// 	}

// 	for i, _ := range t {
// 		if len(t[i]) > 1 {
// 			val := t[i][len(t[i])-1]

// 			for j := len(t[i]) - 1; j >= 1; j-- {
// 				t[i][j] = t[i][j-1]
// 			}

// 			t[i][0] = val
// 		}
// 	}

// 	ans := make([]string, 0, N)
// 	cnt := make([]int, M)
// 	for _, v := range C {
// 		ans = append(ans, t[v-1][cnt[v-1]])
// 		cnt[v-1]++
// 	}

// 	fmt.Println(strings.Join(ans, ""))
// }
