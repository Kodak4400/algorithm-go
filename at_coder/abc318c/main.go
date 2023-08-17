package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, D, P int
	fmt.Fscan(in, &N, &D, &P)

	// 1日周遊パスを買わなかった場合の合計
	sum := 0
	F := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &F[i])
		sum += F[i]
	}

	// 運賃を降順に並び替え
	sort.Sort(sort.Reverse(sort.IntSlice(F)))

	ans := sum
	cnt := 0
	for {
		// 1日周遊パスを購入
		sum += P

		// 1日周遊パスを使用
		for j := 0; j < D; j++ {
			if cnt >= len(F) {
				break
			}
			sum -= F[cnt]
			cnt++
		}

		// 1日周遊パスを使った方がお得な場合、入れ替える
		if ans > sum {
			ans = sum
		}

		if cnt >= len(F) {
			break
		}
	}

	fmt.Fprintln(out, ans)
}
