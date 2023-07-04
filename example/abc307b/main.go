package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	S := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &S[i])
	}

	ex := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f := S[i]
			s := S[j]
			if s == f {
				continue
			}

			ex = chmatch(f + s)
			if ex {
				break
			}
		}
		if ex {
			break
		}
	}

	if ex {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func chmatch(m string) bool {
	for i := 0; i < len(m)/2; i++ {
		if m[i] == m[len(m)-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
