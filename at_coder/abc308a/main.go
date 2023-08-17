package main

import "fmt"

func main() {
	S := make([]int, 8)
	for i := 0; i < 8; i++ {
		fmt.Scanf("%d", &S[i])
	}

	ans1 := false
	for i := 0; i < 7; i++ {
		if S[i] <= S[i+1] {
			ans1 = true
		} else {
			ans1 = false
			break
		}
	}

	ans2 := false
	for i := 0; i < 8; i++ {
		if 100 <= S[i] && S[i] <= 675 {
			ans2 = true
		} else {
			ans2 = false
			break
		}
	}

	ans3 := false
	for i := 0; i < 8; i++ {
		if S[i]%25 == 0 {
			ans3 = true
		} else {
			ans3 = false
			break
		}
	}

	if ans1 && ans2 && ans3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
