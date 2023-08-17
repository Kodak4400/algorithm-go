package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s) // rune型はcode pointを単位として扱う型
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	ss := strings.Split(s, "")

	max := 0
	for i := 0; i < len(ss)-1; i++ {
		for j := len(ss) - 1; j >= i; j-- {
			t := s[i : j+1]
			if t == reverse(t) {
				if max < len(t) {
					max = len(t)
				}
			}
		}
	}

	fmt.Fprintln(out, max)
}
