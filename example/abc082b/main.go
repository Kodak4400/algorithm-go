package main

import (
	"fmt"
	"sort"
	"strings"
)

func Reverse(t []string) []string {
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return t
}

func main() {
	var s string
	_, s_err := fmt.Scan(&s)
	if s_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	var t string
	_, t_err := fmt.Scan(&t)
	if t_err != nil {
		fmt.Println("[Error] input.")
		return
	}

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")

	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	sort.Slice(tt, func(i, j int) bool { return tt[i] < tt[j] })
	ttt := Reverse(tt)
	if strings.Join(ss, "") < strings.Join(ttt, "") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
