package main

import (
	"bufio"
	"fmt"
	"os"
)

func p1() {
	var n, d int64
	fmt.Scan(&n, &d)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	parts1 := []rune(line)

	var count int64
	for i := 0; i < len(parts1); i++ {
		p := parts1[i]
		if p == '.' {
			count++
		}
	}

	ans := count + d

	fmt.Println(ans)
}
