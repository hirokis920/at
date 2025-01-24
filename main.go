package main

import (
	"fmt"
)

func main() {
	var r int
	fmt.Scan(&r)

	x := 0
	ans := 0

	for y := r - 1; y >= 0; y-- {
		for ; inside(x+1, y, r); x++ {
		}
		ans += x

	}

	ans *= 4
	fmt.Println(ans + 1)

}

func inside(x, y, r int) bool {
	x = x*2 + 1
	y = y*2 + 1

	return x*x+y*y <= r*r*4
}
