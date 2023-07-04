package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	C := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &C[i])
	}

	D := make([]string, M+1)
	for i := 1; i <= M; i++ {
		fmt.Scanf("%s", &D[i])
	}

	P := make([]int, M+1)
	for i := 0; i < M+1; i++ {
		fmt.Scanf("%d", &P[i])
	}

	t := 0
	ans := 0
	for i := 1; i <= M; i++ {
		for _, v := range C {
			if D[i] == v {
				ans += P[i]
				t++
			}
		}
	}

	ans += P[0] * (N - t)

	fmt.Println(ans)
}
