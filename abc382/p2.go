package main

import (
	"bufio"
	"fmt"
	"os"
)

func p2() {
	var n, d int64
	fmt.Scan(&n, &d)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	parts1 := []rune(line)

	for i := 0; i < int(d); i++ {
		for j := len(parts1) - 1; j >= 0; j-- {
			p := parts1[j]
			if p == '@' {
				parts1[j] = '.'
				break
			}
		}
	}
	for i := 0; i < len(parts1); i++ {
		fmt.Print(string(parts1[i]))
	}

}
