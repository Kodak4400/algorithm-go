package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(x int, y int, h int, w int, line [][]string) {
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			x2 := x + dx
			y2 := y + dy
			if x2 < 0 || y2 < 0 || x2 > w-1 || y2 > h-1 {
				continue
			}
			// fmt.Println("y2, x2", y2, x2)
			if line[y2][x2] == "#" {
				line[y2][x2] = "."
				// fmt.Println("dfs", y2, x2, line)
				dfs(x2, y2, h, w, line)
			}
		}
	}
}

// グラフ理論
// BFS/DFS or UnionFind(AtCoder ACL dsu)
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h)
	fmt.Fscan(in, &w)

	s := make([]string, h)
	line := make([][]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &s[i])
		line[i] = strings.Split(s[i], "")
	}

	// fmt.Println(line)

	// 黒マスを見つけた数
	cnt := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if line[y][x] == "#" {
				cnt++
				line[y][x] = "."
				// fmt.Println("y,x", y, x)
				dfs(x, y, h, w, line)
			}
		}
	}

	fmt.Println(cnt)
}
