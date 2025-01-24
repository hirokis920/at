package main

import (
	"fmt"
)

func p2() {
	var x int
	fmt.Scan(&x)

	dd := 2

	for i := 0; i < 20; i++ {
		x = x / dd
		if x == 1 {
			ans := dd
			fmt.Println(ans)
			break
		}
		dd++
	}

}
