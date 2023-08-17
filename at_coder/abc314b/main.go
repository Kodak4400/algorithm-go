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

	N := scanInt()

	// 人ごとの賭けた目の情報を格納するスライス
	bets := make([][]int, N)
	// 人ごとの賭けた目の個数を格納するスライス
	counts := make([]int, N)

	for i := 0; i < N; i++ {
		C := scanInt()

		counts[i] = C
		bets[i] = make([]int, C)
		for j := 0; j < C; j++ {
			bets[i][j] = scanInt()
		}
	}

	X := scanInt()

	// fmt.Println(N, bets, counts, X)

	// Xに賭けた人たちの中で、賭けた目の個数が最も少ない値を求める
	minCount := 38 // 最大値よりも大きい初期値
	for i := 0; i < N; i++ {
		for _, bet := range bets[i] {
			if bet == X && counts[i] < minCount {
				minCount = counts[i]
			}
		}
	}

	// fmt.Println("minCount", minCount)

	// Xに賭けて、賭けた目の個数がminCountの人たちの番号を出力する
	var winners []int
	for i := 0; i < N; i++ {
		for _, bet := range bets[i] {
			if bet == X && counts[i] == minCount {
				winners = append(winners, i+1)
				break
			}
		}
	}

	// 結果を出力
	fmt.Println(len(winners))
	for _, winner := range winners {
		fmt.Printf("%d ", winner)
	}
}
