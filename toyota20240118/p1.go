package main

import (
	"fmt"
	"strconv"
)

func p1() {
	var s string
	fmt.Scan(&s)
	srune := []rune(s)

	i, _ := strconv.Atoi(string(srune[0]))

	i2, _ := strconv.Atoi(string(srune[2]))

	ans := i * i2

	fmt.Println(ans)

}
