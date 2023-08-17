package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
7 4
6 7 1 2 3 4 5
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N)
	fmt.Fscan(in, &M)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	fmt.Println(N, M, A)

	sort.Ints(A)

	result := -1
	left := 0
	right := len(A) - 1

	for right >= left {
		mid := left + (right-left)/2
		if A[mid] == M {
			result = mid
			break
		} else if A[mid] > M {
			right = mid - 1
		} else if A[mid] < M {
			left = mid + 1
		}
	}
	fmt.Println(result)

}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var N int
// 	fmt.Fscan(in, &N)

// 	F := make([]int, N)
// 	S := make([]int, N)

// 	for i := 0; i < N; i++ {
// 		fmt.Fscan(in, &F[i])
// 		fmt.Fscan(in, &S[i])
// 	}

// 	max := -1
// 	for i := 0; i < N; i++ {
// 		for j := 0; j < N; j++ {
// 			if i == j {
// 				continue
// 			}
// 			// 味が同じ / 違う
// 			if F[i] == F[j] {
// 				manzoku := S[i] + S[j]/2
// 				if max < manzoku {
// 					max = manzoku
// 				}
// 			} else {
// 				manzoku := S[i] + S[j]
// 				if max < manzoku {
// 					max = manzoku
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(max)
// }
