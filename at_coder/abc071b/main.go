package main

import (
	"fmt"
	"sort"
)

func main() {
	var S string
	_, err := fmt.Scan(&S)
	if err != nil {
		fmt.Println("[error] input.")
		return
	}

	//	ans := "abcdefghijklmnopqrstuvwxyz"

	RS := []rune(S)
	rans := []rune("abcdefghijklmnopqrstuvwxyz")

	sort.Slice(RS, func(i, j int) bool { return string(RS[i]) < string(RS[j]) })

	fmt.Println(string(RS), string(rans))

	ok := false
	for i := 0; i < len(rans); i++ {
		ok = false
		for j := 0; j < len(RS); j++ {
			// fmt.Println(string(rans[i]), string(RS[j]))
			if string(rans[i]) == string(RS[j]) {
				ok = true
				break
			}
		}
		if !ok {
			fmt.Println(string(rans[i]))
			break
		}
	}

	if ok {
		fmt.Println("None")
	}

}
