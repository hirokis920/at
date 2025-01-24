package main

import (
	"fmt"
)

func p1() {

	var n int
	var c1, c2, s string
	fmt.Scan(&n, &c1, &c2, &s)
	srune := []rune(s)

	for i := 0; i < n; i++ {
		if srune[i] != []rune(c1)[0] {
			srune[i] = []rune(c2)[0]
		}
		fmt.Print(string(srune[i]))
	}

}
